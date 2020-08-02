import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import { Button, Segment, Message } from 'semantic-ui-react';
import { makeAuthedBackendRequest } from '../lib/backend';

const JudgeDashboard: React.FC<{ name: string }> = ({ name }) => {
  const [success, setSuccess] = useState<string>('');
  const [failure, setFailure] = useState<string>('');
  const [numEvaluations, setNumEvaluations] = useState<number>(-1);

  const getNumberOfEvaluations = async () => {
    const response = await makeAuthedBackendRequest('get', 'v1/evaluations');
    setNumEvaluations(
      response.data.data.filter((item) => item.CreatedAt === item.UpdatedAt)
        .length,
    );
  };

  useEffect(() => {
    getNumberOfEvaluations();
  }, []);

  return (
    <Segment basic textAlign="left" style={{ padding: '1.5em 2em' }}>
      <p style={{ fontSize: '2.5em', marginBottom: '0' }}>Good morning,</p>
      <p style={{ fontSize: '2.5em', fontWeight: 'bold' }}>{name}</p>
      <p style={{ fontSize: '2em' }}>
        You have <b>{numEvaluations}</b> teams left to judge.
      </p>

      <Link href="/evaluations">
        <Button
          as="a"
          size="big"
          style={{ display: 'block', margin: '1em auto', width: '100%' }}
          color="teal"
        >
          View Assigned Teams
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

export default JudgeDashboard;
