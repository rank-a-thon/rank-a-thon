import React, { useState } from 'react';
import type { NextPage } from 'next';
import { Segment, Table, Menu, Button } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import schedule from '../data/schedule.json';

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
  return (
    <MobilePostAuthContainer title="Schedule" requireAuth>
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
            {schedule.map((timeslot) => (
              <Table.Row disabled={timeslot.over}>
                <Table.Cell>{timeslot.time}</Table.Cell>
                <Table.Cell>{timeslot.activity}</Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default RankingLayout;
