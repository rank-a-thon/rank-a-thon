import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import { Segment, Card } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { makeAuthedBackendRequest } from '../lib/backend';

type PageProps = {
  getWidth?: () => number;
};

type Submission = {
  projId: number;
  projName: string;
  projDesc: string;
  projCoverImg: string;
  teamName: string;
};
const ExploreLayout: NextPage<PageProps> = () => {
  const [submissions, setSubmissions] = useState<Submission[] | null>(null);

  const loadSubmissions = async () => {
    try {
      const submissionsResponse = await makeAuthedBackendRequest(
        'get',
        'v1/submissions/testevent',
      );
      const allSubmissions = submissionsResponse.data.data.map((submission) => {
        return {
          projId: submission.ID,
          projName: submission.project_name,
          projDesc: submission.description,
          projCoverImg: submission.images,
          teamName: submission.team.team_name,
        };
      });
      console.log(allSubmissions);
      setSubmissions(allSubmissions);
    } catch (err) {
      console.error(err.response);
    }
  };

  useEffect(() => {
    loadSubmissions();
  }, []);

  function renderSubmission(submission: Submission) {
    const { projId, projName, projDesc, projCoverImg, teamName } = submission;
    return (
      <Card
        key={projId}
        image={projCoverImg}
        header={projName}
        meta={`by ${teamName}`}
        description={projDesc}
        fluid
      />
    );
  }

  return (
    <MobilePostAuthContainer title="Explore">
      <Segment
        basic
        textAlign="left"
        style={{ padding: '1.5em 2em 0.8em 2em' }}
      >
        {submissions?.map(renderSubmission)}
        <Card
          image="/img/pepekip.png"
          header="PepeMudKip"
          meta="by Mudkip Lovers"
          description="Difficulty choosing your starter pokemon? PepeMudkip is a new-age full-stack PWA that uses modern natural language processing techniques for text-generation to perform named entity recognition on geographical noise data to help you determine the best starter pokemon to choose."
          fluid
        />
        <Card
          image="/img/hammer.jpeg"
          header="ElectionMaster"
          meta="by MerryGandering"
          description="A handshake 🤝 is basically a 🙏promise, a commitment 💍, a tall order 🍾, means I must meet 🏃‍♂️that tall order , and it's for YOU ☝️! And it's for you, in sense that three👌 fingers also pointing me 🙆‍♂️, it's also for ME! 👇 It's for us. 👨‍👩‍👧‍👧And if the result is good 👍👍👍, THUMBS UP MAAAAN... 👍👍👍 and if the result is lousy 😤😤 wad happen? 🧐𝔁𝓾𝓮🥶𝓱𝓾𝓪🧚‍♀️𝓹𝓲𝓪𝓸😻𝓹𝓲𝓪𝓸🗿𝓫𝓮𝓲👺𝓯𝓮𝓷𝓰🤩𝔁𝓲𝓪𝓸😼𝔁𝓲𝓪𝓸👣"
          fluid
        />
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default ExploreLayout;
