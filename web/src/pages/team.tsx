import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Router from 'next/router';
import {
  Segment,
  Input,
  Button,
  List,
  Message,
  Table,
} from 'semantic-ui-react';
import { makeAuthedBackendRequest, makeBackendRequest } from '../lib/backend';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { getMe } from '../data/me';

type PageProps = {
  getWidth?: () => number;
};

type TeamInvite = {
  teamId: number;
  teamName: string;
  inviteSender: string;
};

const DashboardLayout: NextPage<PageProps> = () => {
  const [teamName, setTeamName] = useState<string | null>(null);
  const [teamMembers, setTeamMembers] = useState<string[]>(['']);
  const [createTeamName, setCreateTeamName] = useState<string | null>(null);
  const [inviteEmail, setInviteEmail] = useState<string | null>(null);
  const [inviteSuccess, setInviteSuccess] = useState<boolean>(false);
  const [inviteErr, setInviteErr] = useState<string | null>(null);
  const [routineTask, setRoutineTask] = useState<any>();

  const [teamInvites, setTeamInvites] = useState<TeamInvite[]>(null);

  async function getTeamNameAndMembers() {
    if (!getMe()) {
      return null;
    }
    let response;
    try {
      response = await makeAuthedBackendRequest('get', 'v1/user');
    } catch (err) {
      console.log(err);
    }
    const newMe = response.data.user;
    const allTeamIds = JSON.parse(newMe.team_id_for_event);
    if (Object.keys(allTeamIds).length === 0) {
      return null;
    }
    const teamId = parseInt(allTeamIds.testevent);
    if (teamId === 0) {
      return null;
    }

    try {
      response = await makeAuthedBackendRequest('get', 'v1/team/testevent');
    } catch (err) {
      console.error(err);
    }
    const teamName = response.data.data.team_name;
    console.log(response.data);
    const teamMemberIds = JSON.parse(response.data.data.user_ids);

    const teamMemberNames = await Promise.all(
      teamMemberIds.map(async (id) => {
        try {
          response = await makeAuthedBackendRequest(
            'get',
            `v1/user?userid=${id}`,
          );
          return response.data.user.name;
        } catch (err) {
          console.log(err);
          return 'A Fun Friend';
        }
      }),
    );

    setTeamName(teamName);
    setTeamMembers(teamMemberNames as string[]);
    setTimeout(getTeamNameAndMembers, 5000);
  }

  async function sendMakeTeam() {
    if (createTeamName && createTeamName !== '') {
      await makeAuthedBackendRequest('post', 'v1/team/testevent', {
        team_name: createTeamName,
        is_freshman_team: true,
        is_pre_university_team: false,
        is_beginner_team: false,
      }).then(() => {
        Router.reload();
      });
    }
  }

  async function leaveTeam() {
    const me = getMe();
    try {
      const response = await makeAuthedBackendRequest(
        'delete',
        `v1/remove-team-member/testevent?delete-user-id=${me.ID}`,
      );
      setTeamName(null);
    } catch (err) {
      console.log(err);
    }
  }

  async function sendInvite() {
    setInviteSuccess(false);
    setInviteErr(null);
    const me = getMe();
    try {
      const response = await makeAuthedBackendRequest(
        'post',
        `v1/team-invite`,
        {
          event: 'testevent',
          email: inviteEmail,
        },
      );
      setInviteEmail(null);
      setInviteSuccess(true);
    } catch (err) {
      setInviteErr(err.response.data.message);
    }
  }

  async function getInvites() {
    if (routineTask) {
      clearTimeout(routineTask);
    }
    try {
      const response = await makeAuthedBackendRequest('get', 'v1/team-invites');
      setTeamInvites(
        await Promise.all(
          response.data.data.map(async (inviteResponse) => {
            return {
              teamId: inviteResponse.team_id,
              teamName: (
                await makeAuthedBackendRequest(
                  'get',
                  `v1/team?teamid=${inviteResponse.team_id}`,
                )
              ).data.data.team_name,
              inviteSender: (
                await makeAuthedBackendRequest(
                  'get',
                  `v1/user?userid=${inviteResponse.user_id}`,
                )
              ).data.user.name,
            };
          }),
        ),
      );
    } catch (err) {
      console.error(err.response);
    }
    setRoutineTask(setTimeout(getInvites, 5000));
  }

  useEffect(() => {
    clearTimeout();
    getTeamNameAndMembers();
    getInvites();
  }, []);

  const renderTeamInvites = (invite) => {
    console.log(invite);
    const { teamId, teamName, inviteSender } = invite;
    const sendAccept = () => {
      try {
        makeAuthedBackendRequest(
          'delete',
          `v1/team-invite/accept?teamid=${teamId}`,
        ).then(() => Router.reload());
      } catch (err) {
        console.error(err.response);
      }
    };

    const sendDecline = () => {
      makeAuthedBackendRequest(
        'delete',
        `v1/team-invite/decline?teamid=${teamId}`,
      ).then(getInvites);
    };

    return (
      <Table.Row key={inviteSender}>
        <Table.Cell>
          <p style={{ fontWeight: 'normal', margin: 0 }}>
            <span style={{ fontWeight: 'bold' }}>Team Name:</span> {teamName}
          </p>
          <p style={{ fontWeight: 'normal', margin: '0 0 1em 0' }}>
            <span style={{ fontWeight: 'bold' }}>Sent by:</span> {inviteSender}
          </p>
          <div style={{ textAlign: 'center' }}>
            <Button.Group>
              <Button positive onClick={sendAccept}>
                Accept
              </Button>
              <Button.Or />
              <Button secondary onClick={sendDecline}>
                Decline
              </Button>
            </Button.Group>
          </div>
        </Table.Cell>
      </Table.Row>
    );
  };

  return (
    <MobilePostAuthContainer title="Team" requireAuth>
      {teamName && (
        <Segment
          basic
          textAlign="left"
          style={{ padding: '1.5em 2em 0.8em 2em' }}
        >
          <p style={{ fontSize: '1.4em', marginBottom: '0.4em' }}>
            <span style={{ fontWeight: 'bold' }}>Team Name:</span> {teamName}
          </p>
          <p style={{ fontSize: '1.4em', margin: '0', fontWeight: 'bold' }}>
            Members:
          </p>
          <List style={{ fontSize: '1.4em', marginTop: '0.4em' }}>
            {teamMembers.map((name) => {
              return (
                <List.Item key={name}>
                  <List.Icon name="user circle" />
                  <List.Content>{name}</List.Content>
                </List.Item>
              );
            })}
          </List>

          <Segment color="violet" style={{ overflow: 'hidden' }}>
            <p
              style={{
                fontSize: '1.1em',
                fontWeight: 'bold',
                margin: '0.2em 0em',
              }}
            >
              Invite Members
            </p>
            <p
              style={{
                margin: '0.2em 0em',
              }}
            >
              You can invite {4 - teamMembers.length} more member
              {4 - teamMembers.length === 1 ? '' : 's'}.
            </p>
            <Input
              placeholder="Member Email Address"
              style={{ margin: '0.4em 0', width: '100%' }}
              onChange={(e) => setInviteEmail(e.target.value)}
            />
            <Button
              primary
              disabled={4 - teamMembers.length <= 0}
              size="medium"
              onClick={sendInvite}
              style={{
                alignItems: 'flex-end',
                margin: '0.8em 0 0 0',
              }}
            >
              Invite
            </Button>
            {inviteErr && (
              <Message negative style={{ margin: '0.8em 0 0 0' }}>
                {inviteErr}
              </Message>
            )}
            {inviteSuccess && (
              <Message positive style={{ margin: '0.8em 0 0 0' }}>
                Successfully sent an invite. Ask your friend to accept it!
              </Message>
            )}
          </Segment>

          <Message warning>
            <Message.Header>
              Your team has not made your project submission!
            </Message.Header>
            <p>
              Please ensure you submit your project before the judging deadline
              of 0900hrs on Day 2 of the event!
            </p>
          </Message>

          <Button
            size="medium"
            style={{ display: 'block', margin: '1em auto', width: '100%' }}
            primary
            onClick={() => alert('Not supported yet. Sorry!')}
          >
            Submit your project
          </Button>
          <Button
            size="medium"
            style={{ display: 'block', margin: '1em auto', width: '100%' }}
            onClick={leaveTeam}
            negative
          >
            Leave Team
          </Button>
        </Segment>
      )}

      {!teamName && (
        <Segment
          basic
          textAlign="left"
          style={{ padding: '1.5em 2em 0.8em 2em' }}
        >
          <p style={{ fontSize: '1.4em', margin: '0' }}>
            You have not joined a team yet.
          </p>
          <Segment color="violet" style={{ overflow: 'hidden' }}>
            <p
              style={{
                fontSize: '1.1em',
                fontWeight: 'bold',
                margin: '0.2em 0em',
              }}
            >
              Create a Team
            </p>
            <Input
              placeholder="Team Name"
              style={{ margin: '0.4em 0', width: '100%' }}
              onChange={(e) => setCreateTeamName(e.target.value)}
            />
            <Button
              primary
              size="medium"
              floated="right"
              onClick={sendMakeTeam}
              style={{ margin: '0.8em 0 0 0' }}
            >
              Submit
            </Button>
          </Segment>
          <Segment color="teal" style={{ overflow: 'hidden' }}>
            <p
              style={{
                fontSize: '1.1em',
                fontWeight: 'bold',
                margin: '0.2em 0em',
              }}
            >
              Accept Team Invites
            </p>
            {teamInvites && teamInvites.length === 0 ? (
              <p
                style={{
                  fontSize: '1em',
                  fontWeight: 'lighter',
                  margin: '0.8em 0em',
                }}
              >
                You do not have any invites at the moment. Did your friend send
                an invite?
              </p>
            ) : (
              <Table celled singleLine>
                <Table.Body>
                  {teamInvites && teamInvites.map(renderTeamInvites)}
                </Table.Body>
              </Table>
            )}
          </Segment>
        </Segment>
      )}
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
