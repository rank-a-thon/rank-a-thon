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
          Rate each category from 1 (worst) to 5 (best):
        </p>

        <div style={{ textAlign: 'center', margin: '1em' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
              marginBottom: '0.4em',
            }}
          >
            Annoyingness
          </p>
          <RatingScale setCategoryScore={console.log} />
        </div>

        <div style={{ textAlign: 'center', margin: '1em' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
              marginBottom: '0.4em',
            }}
          >
            Entertainment
          </p>
          <RatingScale setCategoryScore={console.log} />
        </div>

        <div style={{ textAlign: 'center', margin: '1em' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
              marginBottom: '0.4em',
            }}
          >
            Beauty
          </p>
          <RatingScale setCategoryScore={console.log} />
        </div>

        <div style={{ textAlign: 'center', margin: '1em' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
              marginBottom: '0.4em',
            }}
          >
            Social Usefulness
          </p>
          <RatingScale setCategoryScore={console.log} />
        </div>

        <div style={{ textAlign: 'center', margin: '1em' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
              marginBottom: '0.4em',
            }}
          >
            Hardware
          </p>
          <RatingScale setCategoryScore={console.log} />
        </div>

        <div style={{ textAlign: 'center', margin: '1em' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
              marginBottom: '0.4em',
            }}
          >
            Awesomely Useless
          </p>
          <RatingScale setCategoryScore={console.log} />
        </div>

        <hr />
        <div style={{ textAlign: 'center', margin: '1em' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
              marginBottom: '0.4em',
            }}
          >
            Overall Rating
          </p>
          <RatingScale setCategoryScore={console.log} />
        </div>

        <div style={{ textAlign: 'center', margin: '2em 0' }}>
          <p
            style={{
              fontSize: '1.2em',
              fontWeight: 'bold',
            }}
          >
            All changes are saved automatically.
          </p>
        </div>
      </Segment>
    </MobilePostAuthContainer>
  );
};

const RatingScale: React.FC<{ setCategoryScore: (number) => void }> = ({
  setCategoryScore,
}) => {
  const [selected, setSelected] = useState<number>(0);
  return (
    <Button.Group>
      {[1, 2, 3, 4, 5].map((rating) => {
        const rateOnClick = () => {
          setSelected(rating);
          setCategoryScore(rating);
        };
        return (
          <Button
            key={rating}
            active={rating === selected}
            onClick={rateOnClick}
          >
            {rating}
          </Button>
        );
      })}
    </Button.Group>
  );
};

export default EvaluateLayout;
