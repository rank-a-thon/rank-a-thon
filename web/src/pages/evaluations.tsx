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
      return allSubmissions;
    } catch (err) {
      console.error(err.response);
    }
  };

  const getEvaluations = async () => {
    const response = await makeAuthedBackendRequest('get', 'v1/evaluations');
    const submissionsToJudge = response.data.data.map(
      (evaluation) => evaluation.submission_id,
    );
    const allSubmissions = await loadSubmissions();
    const toJudge = allSubmissions.filter((submission) =>
      submissionsToJudge.includes(submission.projId),
    );
    setSubmissions(toJudge);
  };

  useEffect(() => {
    getEvaluations();
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
    <MobilePostAuthContainer title="Judge">
      <Segment
        basic
        textAlign="left"
        style={{ padding: '1.5em 2em 0.8em 2em' }}
      >
        {submissions?.map(renderSubmission)}
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default ExploreLayout;
