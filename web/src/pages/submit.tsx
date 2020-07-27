import React, { useState, useEffect, useRef } from 'react';
import type { NextPage } from 'next';
import Router from 'next/router';
import {
  Segment,
  Input,
  Button,
  Form,
  Checkbox,
  Card,
  List,
  Message,
  Table,
  Image,
} from 'semantic-ui-react';
import {
  makeAuthedBackendRequest,
  makeBackendRequest,
  sendAuthedFormData,
} from '../lib/backend';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import { getMe } from '../data/me';
import axios from 'axios';

type PageProps = {
  getWidth?: () => number;
};

const uploadFile = async (file) => {
  const formData = new FormData();
  formData.append('file', file);
  const uploadResponse = await sendAuthedFormData(
    'post',
    'v1/submission-file-upload/testevent',
    formData,
  );
  return uploadResponse.data.image_url;
};

const SubmitLayout: NextPage<PageProps> = () => {
  const fileUploadRef = useRef<any>(null);
  const [imageUploading, setimageUploading] = useState<boolean>(false);
  const [coverImageUrl, setCoverImageUrl] = useState<string>('');
  const [coverImageName, setCoverImageName] = useState<string>('');
  const [error, setError] = useState<string>('');
  const [success, setSuccess] = useState<string>('');

  const [projName, setProjName] = useState<string>('');
  const [projDesc, setProjDesc] = useState<string>('');
  const [rulesChecked, setRulesChecked] = useState<boolean>(false);
  const [voidChecked, setVoidChecked] = useState<boolean>(false);
  const [submitSending, setSubmitSending] = useState<boolean>(false);

  const [hasExistingSubmissions, setHasExistingSubmissions] = useState<boolean>(
    false,
  );

  const uploadCoverImage = async (event) => {
    setimageUploading(true);
    const file = event.target.files[0];
    const uploadedImageUrl = await uploadFile(file);
    setCoverImageUrl(uploadedImageUrl);
    setCoverImageName(file.name);
    setimageUploading(false);
  };

  const sendSubmission = async () => {
    setSubmitSending(true);
    setError('');
    setSuccess('');
    if (
      !coverImageUrl ||
      !projName ||
      !projDesc ||
      !rulesChecked ||
      !voidChecked
    ) {
      setError('Please fill up all fields!');
      setSubmitSending(false);
      return;
    }

    try {
      await makeAuthedBackendRequest(
        hasExistingSubmissions ? 'put' : 'post',
        'v1/submission/testevent',
        {
          project_name: projName,
          description: projDesc,
          images: [coverImageUrl],
        },
      );
    } catch (err) {
      setError(err.response?.data?.error || JSON.stringify(err.response));
      setSubmitSending(false);
      return;
    }
    setSubmitSending(false);
    setSuccess(
      hasExistingSubmissions
        ? 'Project updated successfully! All the best for judging! 😇'
        : 'Project submitted successfully! All the best for judging! 😇',
    );
    setHasExistingSubmissions(true);
  };

  const loadPreviousSubmission = async () => {
    try {
      const prevSubmissionResponse = await makeAuthedBackendRequest(
        'get',
        'v1/submission/testevent',
      );
      const {
        project_name: projName,
        description: projDesc,
        images: coverImageUrl,
      } = prevSubmissionResponse.data.data;
      setHasExistingSubmissions(true);
      setProjName(projName);
      setProjDesc(projDesc);
      setCoverImageUrl(coverImageUrl);
      setCoverImageName(coverImageUrl.split('/').slice(-1)[0]);
      setVoidChecked(true);
      setRulesChecked(true);
    } catch (err) {
      if (err.response.status === 404) {
        setHasExistingSubmissions(false);
      }
    }
  };

  useEffect(() => {
    loadPreviousSubmission();
  }, []);

  return (
    <MobilePostAuthContainer
      title={hasExistingSubmissions ? 'Update' : 'Submit'}
      requireAuth
    >
      <Segment basic textAlign="left" style={{ padding: '1.5em 2em' }}>
        <p style={{ fontSize: '1.4em' }}>
          {hasExistingSubmissions
            ? "Update your team's submission"
            : 'Make a submission for your team!'}
        </p>
        <Form error={!!error} success={!!success}>
          <Form.Input
            required
            fluid
            label="Project Name"
            value={projName}
            onChange={(e) => setProjName(e.target.value)}
            placeholder="Give your project a name!"
          />
          <Form.TextArea
            required
            style={{ fontFamily: 'Lato, sans-serif' }}
            label="Description"
            value={projDesc}
            onChange={(e) =>
              setProjDesc((e.target as HTMLTextAreaElement).value)
            }
            placeholder="Write a short paragraph about your project!"
            rows={8}
          />
          <Form.Field required>
            <label>Cover Image</label>
            {coverImageUrl && (
              <div
                style={{
                  width: '100%',
                  border: '1px solid #d1d1d1',
                  borderRadius: '10px',
                  padding: '0.5em 0.8em',
                  marginBottom: '1em',
                }}
              >
                <Image inline src={coverImageUrl} size="tiny" />
                <span style={{ marginLeft: '0.8em', color: '#a8a8a8' }}>
                  {coverImageName}
                </span>
              </div>
            )}
            <Button
              loading={imageUploading}
              content={
                coverImageUrl ? 'Replace Cover Image' : 'Upload Cover Image'
              }
              labelPosition="left"
              icon="picture"
              onClick={() => fileUploadRef.current.click()}
            />
            <input
              ref={fileUploadRef}
              type="file"
              accept="image/*"
              hidden
              onChange={uploadCoverImage}
            />
          </Form.Field>

          <Form.Field required>
            <Checkbox
              checked={rulesChecked}
              onChange={() => setRulesChecked(!rulesChecked)}
              label="I acknowledge that my team’s submission adheres to the rules and regulations of the hackathon."
            />
          </Form.Field>
          <Form.Field required>
            <Checkbox
              checked={voidChecked}
              onChange={() => setVoidChecked(!voidChecked)}
              label="I acknowledge that if any part of my team’s submission is found to contravene the rules, or is incomplete, the organisers have the right to void my team’s submission."
            />
          </Form.Field>

          <Message error content={error} />
          <Message success content={success} />
          <Button
            primary
            type="submit"
            onClick={sendSubmission}
            loading={submitSending}
          >
            Submit
          </Button>
        </Form>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default SubmitLayout;
