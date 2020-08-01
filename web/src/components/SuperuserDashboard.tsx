import React from 'react';
import Link from 'next/link';
import { Button, Segment } from 'semantic-ui-react';

const SuperDashboard: React.FC<{ name: string }> = ({ name }) => (
  <Segment basic textAlign="left" style={{ padding: '1.5em 2em' }}>
    <p style={{ fontSize: '2.5em', marginBottom: '0' }}>Hi Superuser,</p>
    <p style={{ fontSize: '2.5em', fontWeight: 'bold' }}>{name}</p>
    <p style={{ fontSize: '2em' }}>
      With great power comes great responsibility.
    </p>

    <Link href="/announcements">
      <Button
        as="a"
        size="big"
        style={{ display: 'block', margin: '1em auto', width: '100%' }}
        color="pink"
      >
        View Announcements
      </Button>
    </Link>
    <Link href="/team">
      <Button
        as="a"
        size="big"
        color="purple"
        style={{ display: 'block', margin: '1em auto', width: '100%' }}
      >
        Edit Team
      </Button>
    </Link>
    <Link href="/explore">
      <Button
        as="a"
        size="big"
        color="violet"
        style={{ display: 'block', margin: '1em auto', width: '100%' }}
      >
        Explore Projects
      </Button>
    </Link>
  </Segment>
);

export default SuperDashboard;
