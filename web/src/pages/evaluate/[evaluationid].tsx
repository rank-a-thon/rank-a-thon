import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Router, { useRouter } from 'next/router';
import {
  Segment,
  Input,
  Button,
  List,
  Message,
  Table,
} from 'semantic-ui-react';
import {
  makeAuthedBackendRequest,
  makeBackendRequest,
} from '../../lib/backend';
import MobilePostAuthContainer from '../../components/MobilePostAuthContainer';
import { getMe } from '../../data/me';
import Link from 'next/link';

type PageProps = {
  getWidth?: () => number;
};

const EvaluateLayout: NextPage<PageProps> = () => {
  const router = useRouter();
  const { evaluationid } = router.query;

  return (
    <MobilePostAuthContainer title="Evaluate" requireAuth>
      <Segment
        basic
        textAlign="left"
        style={{ padding: '1.5em 2em 0.8em 2em' }}
      >
        <p style={{ fontSize: '1.4em', margin: '0' }}>
          You have not joined a team yet. {evaluationid}
        </p>
        <Segment color="violet" style={{ overflow: 'hidden' }}>
          <p
            style={{
              fontSize: '1.1em',
              fontWeight: 'bold',
              margin: '0.2em 0em',
            }}
          >
            Create a Team
          </p>
          <Input
            placeholder="Team Name"
            style={{ margin: '0.4em 0', width: '100%' }}
            onChange={null}
          />
          <Button
            primary
            size="medium"
            floated="right"
            onClick={null}
            style={{ margin: '0.8em 0 0 0' }}
          >
            Submit
          </Button>
        </Segment>
        <Segment color="teal" style={{ overflow: 'hidden' }}>
          <p
            style={{
              fontSize: '1.1em',
              fontWeight: 'bold',
              margin: '0.2em 0em',
            }}
          >
            Accept Team Invites
          </p>
        </Segment>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default EvaluateLayout;
