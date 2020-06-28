import React from 'react';
import type { NextPage } from 'next';
import { Segment } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  return (
    <MobilePostAuthContainer title="Explore">
      <Segment
        basic
        textAlign="left"
        style={{ padding: '1.5em 2em 0.8em 2em' }}
      >
        <p style={{ fontSize: '1.4em', margin: '0' }}>
          View other team's projects:
        </p>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
