import React, { useState } from 'react';
import Link from 'next/link';
import { Button, Segment, Message } from 'semantic-ui-react';
import { makeAuthedBackendRequest } from '../lib/backend';

const SuperDashboard: React.FC<{ name: string }> = ({ name }) => {
  const [success, setSuccess] = useState<string>('');
  const [failure, setFailure] = useState<string>('');

  const sendStartEval = async () => {
    setSuccess('');
    setFailure('');
    console.log('wtf');
    try {
      await makeAuthedBackendRequest(
        'put',
        'v1/ranker/start-evaluations/testevent',
      );
      setSuccess(
        'Request sent successfully! Evaluations are now being generated and assigned to judges.',
      );
    } catch (err) {
      console.log(err);
      if (err.response.status === 400) {
        setSuccess(
          'Request sent successfully! Evaluations are now being generated and assigned to judges.',
        );
      } else {
        setFailure(err?.response?.data?.message || err);
      }
    }
  };

  const sendGenerateRankings = async () => {
    setSuccess('');
    setFailure('');
    try {
      await makeAuthedBackendRequest(
        'put',
        'v1/ranker/calculate-team-rankings/testevent',
      );
      setSuccess(
        'Request sent successfully! Rankings are now being generated, and might take a while...',
      );
    } catch (err) {
      setFailure(err?.response?.data?.message || err);
    }
  };

  return (
    <Segment basic textAlign="left" style={{ padding: '1.5em 2em' }}>
      <p style={{ fontSize: '2.5em', marginBottom: '0' }}>Hi Superuser,</p>
      <p style={{ fontSize: '2.5em', fontWeight: 'bold' }}>{name}</p>
      <p style={{ fontSize: '2em' }}>
        With great power comes great responsibility.
      </p>

      <Button
        as="a"
        size="big"
        style={{ display: 'block', margin: '1em auto', width: '100%' }}
        color="orange"
        onClick={sendStartEval}
      >
        Start Evaluations
      </Button>
      <Button
        as="a"
        size="big"
        color="teal"
        style={{ display: 'block', margin: '1em auto', width: '100%' }}
        onClick={sendGenerateRankings}
      >
        Generate Rankings
      </Button>
      <Link href="/explore">
        <Button
          as="a"
          size="big"
          color="olive"
          style={{ display: 'block', margin: '1em auto', width: '100%' }}
        >
          View Generated Rankings
        </Button>
      </Link>
      {failure && (
        <Message negative style={{ margin: '0.8em 0 0 0' }}>
          {failure}
        </Message>
      )}
      {success && (
        <Message positive style={{ margin: '0.8em 0 0 0' }}>
          {success}
        </Message>
      )}
    </Segment>
  );
};

export default SuperDashboard;
