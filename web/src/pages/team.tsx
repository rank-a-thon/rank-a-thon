import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Router from 'next/router';
import { Segment, Input, Button } from 'semantic-ui-react';
import { makeAuthedBackendRequest } from '../lib/backend';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  const [teamName, setTeamName] = useState<string | null>(null);
  const [createTeamName, setCreateTeamName] = useState<string | null>(null);

  async function getTeamName() {
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

    try {
      response = await makeAuthedBackendRequest('get', 'v1/team/testevent');
    } catch (err) {
      console.log(err);
    }
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

  useEffect(() => {
    getTeamName().then((teamName) => setTeamName(teamName));
  }, []);
  return (
    <MobilePostAuthContainer title="Team" requireAuth>
      {teamName && (
        <Segment
          basic
          textAlign="left"
          style={{ padding: '1.5em 2em 0.8em 2em' }}
        >
          <p style={{ fontSize: '1.4em', margin: '0' }}>
            Manage your team {teamName}
          </p>
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
