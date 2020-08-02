import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import { useRouter } from 'next/router';
import { Segment, Button } from 'semantic-ui-react';
import { makeAuthedBackendRequest } from '../../lib/backend';
import MobilePostAuthEvaluateContainer from '../../components/MobilePostAuthEvaluateContainer';

type PageProps = {
  getWidth?: () => number;
};

const EvaluateLayout: NextPage<PageProps> = () => {
  const router = useRouter();
  const { evaluationid } = router.query;
  console.log(evaluationid);

  const [annoying, setAnnoying] = useState<number>(0);
  const [entertaining, setEntertaining] = useState<number>(0);
  const [beauty, setBeauty] = useState<number>(0);
  const [useful, setUseful] = useState<number>(0);
  const [hardware, setHardware] = useState<number>(0);
  const [useless, setUseless] = useState<number>(0);
  const [overall, setOverall] = useState<number>(0);
  const [evalFetched, setEvalFetched] = useState<boolean>(false);

  const getEvaluation = async () => {
    const response = await makeAuthedBackendRequest(
      'get',
      `v1/evaluation/${evaluationid}`,
    );
    const {
      main_rating,
      annoying_rating,
      entertaining_rating,
      beautiful_rating,
      socially_useful_rating,
      hardware_rating,
      awesomely_useless_rating,
    } = response?.data?.data;
    setOverall(main_rating);
    setAnnoying(annoying_rating);
    setEntertaining(entertaining_rating);
    setBeauty(beautiful_rating);
    setUseful(socially_useful_rating);
    setHardware(hardware_rating);
    setUseless(awesomely_useless_rating);
    setEvalFetched(true);
    console.log('fetched');
    console.log(response);
  };

  useEffect(() => {
    if (evaluationid !== undefined) {
      getEvaluation();
    }
  }, [evaluationid]);

  const updateEvaluation = async () => {
    try {
      await makeAuthedBackendRequest('put', `v1/evaluation/${evaluationid}`, {
        main_rating: overall,
        annoying_rating: annoying,
        entertaining_rating: entertaining,
        beautiful_rating: beauty,
        socially_useful_rating: useful,
        hardware_rating: hardware,
        awesomely_useless_rating: useless,
      });
    } catch (err) {
      console.error(err.response);
    }
  };

  useEffect(() => {
    if (
      evalFetched &&
      ![
        overall,
        annoying,
        entertaining,
        beauty,
        useful,
        hardware,
        useless,
      ].includes(0)
    ) {
      updateEvaluation();
    }
  }, [overall, annoying, entertaining, beauty, useful, hardware, useless]);

  return (
    <MobilePostAuthEvaluateContainer title="Evaluate" requireAuth>
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
          <RatingScale setCategoryScore={setAnnoying} selected={annoying} />
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
          <RatingScale
            setCategoryScore={setEntertaining}
            selected={entertaining}
          />
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
          <RatingScale setCategoryScore={setBeauty} selected={beauty} />
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
          <RatingScale setCategoryScore={setUseful} selected={useful} />
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
          <RatingScale setCategoryScore={setHardware} selected={hardware} />
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
          <RatingScale setCategoryScore={setUseless} selected={useless} />
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
          <RatingScale setCategoryScore={setOverall} selected={overall} />
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
    </MobilePostAuthEvaluateContainer>
  );
};

const RatingScale: React.FC<{
  setCategoryScore: (number) => void;
  selected: number;
}> = ({ setCategoryScore, selected }) => {
  return (
    <Button.Group>
      {[1, 2, 3, 4, 5].map((rating) => {
        const rateOnClick = () => {
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
