import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import { Segment, Table, Button } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { makeAuthedBackendRequest } from '../lib/backend';

type PageProps = {
  getWidth?: () => number;
};
const prizes = [
  { name: 'Top 8', key: 'main' },
  { name: 'Most Annoying', key: 'annoying' },
  { name: 'Most Entertaining', key: 'entertaining' },
  { name: 'Most Beautiful', key: 'beautiful' },
  { name: 'Most Socially Useful', key: 'socially_useful' },
  { name: 'Best Hardware', key: 'hardware' },
  { name: 'Most Awesomely Useless', key: 'awesomely_useless' },
];

const RankingLayout: NextPage<PageProps> = () => {
  const [currCat, setCurrCat] = useState<number>(0);
  const [winners, setWinners] = useState<any>(null);

  const loadSubmissions = async () => {
    try {
      const submissionsResponse = await makeAuthedBackendRequest(
        'get',
        'v1/submissions/testevent',
      );
      const allSubmissions = submissionsResponse?.data?.data?.map(
        (submission) => {
          return {
            projId: submission.ID,
            projName: submission.project_name,
            projDesc: submission.description,
            projCoverImg: submission.images,
            teamName: submission.team.team_name,
          };
        },
      );
      return allSubmissions;
    } catch (err) {
      console.error(err.response);
    }
  };

  const getCategoryWinners = async () => {
    try {
      const response = await makeAuthedBackendRequest(
        'post',
        'v1/ranker/team-rankings-by-range/testevent',
        {
          category: prizes[currCat].key,
          start_index: 0,
          end_index: 8,
        },
      );
      const winnersSubmissions = response?.data?.data?.map(
        (winner) => winner.submission_id,
      );
      const allSubmissions = await loadSubmissions();
      const winners = winnersSubmissions
        ?.map((subId) => {
          return allSubmissions.filter((sub) => sub.projId == subId)[0]
            .teamName;
        })
        .filter((v, i, s) => s.indexOf(v) === i);
      console.log(winners);
      setWinners(winners);
    } catch (err) {
      console.error(err.response);
    }
  };

  useEffect(() => {
    getCategoryWinners();
  }, [currCat]);

  return (
    <MobilePostAuthContainer title="Winners" requireAuth>
      <Segment basic textAlign="center" style={{ padding: '1em' }}>
        <p style={{ fontSize: '1.3em', marginBottom: '0' }}>
          <b>Cur:</b> {prizes[currCat].name}
        </p>
        <p style={{ fontSize: '1.3em', marginBottom: '0' }}>
          <b>Next:</b> {prizes[currCat + 1]?.name}
        </p>
        <Button.Group>
          <Button onClick={() => setCurrCat(Math.max(0, currCat - 1))}>
            Prev
          </Button>
          <Button
            primary
            onClick={() => setCurrCat(Math.min(prizes.length - 1, currCat + 1))}
          >
            Next
          </Button>
        </Button.Group>
        <Table
          celled
          unstackable
          selectable
          striped
          color="violet"
          inverted
          textAlign="center"
        >
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Pos</Table.HeaderCell>
              <Table.HeaderCell>Team Name</Table.HeaderCell>
            </Table.Row>
          </Table.Header>

          <Table.Body>
            {winners?.map((name, idx) => (
              <Table.Row>
                <Table.Cell>{idx + 1}</Table.Cell>
                <Table.Cell>{name}</Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default RankingLayout;
