import React from 'react';
import type { NextPage } from 'next';
import { Segment } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  return (
    <MobilePostAuthContainer title="Schedule">
      <Segment
        basic
        textAlign="left"
        style={{ padding: '1.5em 2em 0.8em 2em' }}
      >
        <p style={{ fontSize: '1.4em', margin: '0' }}>We are located at:</p>
        <p style={{ fontSize: '1.4em', fontWeight: 'bolder' }}>
          Cinnamon-Tembusu Dining Hall
        </p>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
