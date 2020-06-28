import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Link from 'next/link';
import { Button, Segment } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { getMe } from '../data/me';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  const [name, setName] = useState<string>('');
  useEffect(() => {
    const me = getMe();
    setName(me ? me.name : 'Christopher Goh');
  }, []);

  return (
    <MobilePostAuthContainer title="Dashboard" requireAuth>
      <Segment basic textAlign="left" style={{ padding: '1.5em 2em' }}>
        <p style={{ fontSize: '2.5em', marginBottom: '0' }}>Good morning,</p>
        <p style={{ fontSize: '2.5em', fontWeight: 'bold' }}>{name}</p>
        <p style={{ fontSize: '2em' }}>What would you like to do today?</p>

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
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
