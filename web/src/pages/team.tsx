import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Router from 'next/router';
import { Segment, Input, Button, List, Message } from 'semantic-ui-react';
import { makeAuthedBackendRequest } from '../lib/backend';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { getMe } from '../data/me';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  const [teamName, setTeamName] = useState<string | null>(null);
  const [createTeamName, setCreateTeamName] = useState<string | null>(null);

  async function getTeamNameAndMembers() {
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
      console.log(err);
    }
    const teamName = response.data.data.team_name;
    const teamMemberIds = JSON.parse(response.data.data.user_ids);
    // TODO: make this return team member names as well
    return response.data.data.team_name;
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

  useEffect(() => {
    getTeamNameAndMembers().then((teamName) => setTeamName(teamName));
  }, []);
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
            <List.Item>
              <List.Icon name="user circle" />
              <List.Content>{getMe().name}</List.Content>
            </List.Item>
            <List.Item>
              <List.Icon name="user circle" />
              <List.Content>Len Beong</List.Content>
            </List.Item>
            <List.Item>
              <List.Icon name="user circle" />
              <List.Content>Soo Juen Yien</List.Content>
            </List.Item>
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
              You can invite 1 more member.
            </p>
            <Input
              placeholder="Member Email Address"
              style={{ margin: '0.4em 0', width: '100%' }}
              onChange={(e) => setCreateTeamName(e.target.value)}
            />
            <Button
              primary
              size="medium"
              floated="right"
              onClick={() => alert('Not supported yet. Sorry!')}
              style={{ margin: '0.8em 0 0 0' }}
            >
              Invite
            </Button>
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
            <p
              style={{
                fontSize: '1em',
                fontWeight: 'lighter',
                margin: '0.8em 0em',
              }}
            >
              You do not have any invites at the moment. Did your friend send an
              invite?
            </p>
          </Segment>
        </Segment>
      )}
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
