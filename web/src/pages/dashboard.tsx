import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Link from 'next/link';
import { Button, Segment } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { getMe } from '../data/me';
import { makeAuthedBackendRequest } from '../lib/backend';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  const [name, setName] = useState<string>('');
  const [userType, setUserType] = useState<'user' | 'judge' | 'superuser'>(
    'user',
  );

  const getNameOnLoad = async () => {
    const response = await makeAuthedBackendRequest('get', 'v1/user');
    setName(response?.data?.user?.name);
  };

  const getUserTypeOnLoad = async () => {
    const response = await makeAuthedBackendRequest('get', 'v1/user');
    switch (response.data.user.user_type) {
      case 2:
        setUserType('judge');
        return;
      case 3:
        setUserType('superuser');
        return;
      default:
        setUserType('user');
        return;
    }
  };

  useEffect(() => {
    getNameOnLoad();
    getUserTypeOnLoad();
  }, []);

  return (
    <MobilePostAuthContainer
      title="Dashboard"
      requireAuth
      judge={userType === 'judge'}
      superuser={userType === 'superuser'}
    >
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
