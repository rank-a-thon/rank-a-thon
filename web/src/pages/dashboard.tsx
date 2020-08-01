import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { makeAuthedBackendRequest } from '../lib/backend';
import UserDashboard from '../components/UserDashboard';

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
      <UserDashboard name={name} />
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
