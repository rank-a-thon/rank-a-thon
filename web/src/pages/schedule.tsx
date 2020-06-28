import React from 'react';
import type { NextPage } from 'next';
import { Segment, Table } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import schedule from '../data/schedule.json';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  return (
    <MobilePostAuthContainer title="Schedule" requireAuth>
      <Segment basic textAlign="center" style={{ padding: '1em' }}>
        <p style={{ fontSize: '1.3em', marginBottom: '0' }}>
          This schedule will always be updated:
        </p>
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
              <Table.HeaderCell>Time</Table.HeaderCell>
              <Table.HeaderCell>Activity</Table.HeaderCell>
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

export default DashboardLayout;
