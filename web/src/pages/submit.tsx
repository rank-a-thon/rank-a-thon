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

  const uploadCoverImage = async (event) => {
    const file = event.target.files[0];
    const uploadedImageUrl = await uploadFile(file);
    setCoverImageUrl(uploadedImageUrl);
  };

  return (
    <MobilePostAuthContainer title="Submit" requireAuth>
      <Segment basic textAlign="left" style={{ padding: '1.5em 2em' }}>
        <p style={{ fontSize: '1.4em' }}>Make a submission for your team!</p>
        <Form>
          <Form.Field>
            <Form.Input
              fluid
              label="Project Name"
              placeholder="Give your project a name!"
            />
          </Form.Field>
          <Form.Field>
            <Form.TextArea
              label="Description"
              placeholder="Write a short paragraph about your project!"
              rows={8}
            />
          </Form.Field>
          <Form.Field>
            <Button
              content="Upload Cover Image"
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
            <Card>
              <Image inline src={coverImageUrl} size="tiny" />
              Wow
            </Card>
          )}
          <Form.Field>
            <Checkbox label="I acknowledge that my team’s submission adheres to the rules of regulations of the hackathon, and if any part of my team’s submission is found to contravene the rules, or is incomplete, the organisers have the right to void my team’s submission." />
          </Form.Field>
          <Button primary type="submit">
            Submit
          </Button>
        </Form>
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default SubmitLayout;
