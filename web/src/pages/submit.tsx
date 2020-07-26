import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Router from 'next/router';
import {
  Segment,
  Input,
  Button,
  List,
  Message,
  Table,
} from 'semantic-ui-react';
import { makeAuthedBackendRequest, makeBackendRequest } from '../lib/backend';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { getMe } from '../data/me';

type PageProps = {
  getWidth?: () => number;
};


const SubmitLayout: NextPage<PageProps> = () => {

  return (
    <MobilePostAuthContainer title="Submit" requireAuth>
      WTF
    </MobilePostAuthContainer>
  );
};

export default SubmitLayout;
