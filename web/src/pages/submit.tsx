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
  const [coverImageUrl, setCoverImageUrl] = useState<string>('');
  const [coverImageName, setCoverImageName] = useState<string>('');
  const [error, setError] = useState<string>('');
  const [success, setSuccess] = useState<string>('');

  const [projName, setProjName] = useState<string>('');
  const [projDesc, setProjDesc] = useState<string>('');
  const [rulesChecked, setRulesChecked] = useState<boolean>(false);
  const [voidChecked, setVoidChecked] = useState<boolean>(false);

  const uploadCoverImage = async (event) => {
    const file = event.target.files[0];
    const uploadedImageUrl = await uploadFile(file);
    setCoverImageUrl(uploadedImageUrl);
    setCoverImageName(file.name);
  };

  return (
    <MobilePostAuthContainer title="Submit" requireAuth>
      <Segment basic textAlign="left" style={{ padding: '1.5em 2em' }}>
        <p style={{ fontSize: '1.4em' }}>Make a submission for your team!</p>
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
            <Button
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
          <Button primary type="submit">
            Submit
          </Button>
        </Form>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default SubmitLayout;
