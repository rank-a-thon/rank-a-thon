import React from 'react';
import type { NextPage } from 'next';
import { Segment } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  return (
    <MobilePostAuthContainer title="Map">
      <Segment
        basic
        textAlign="left"
        style={{ padding: '1.5em 2em 0.8em 2em' }}
      >
        <p style={{ fontSize: '1.4em', margin: '0' }}>We are located at:</p>
        <p style={{ fontSize: '1.4em', fontWeight: 'bolder' }}>
          Cinnamon-Tembusu Dining Hall
        </p>
      </Segment>
      <div style={{ height: 'calc(100vh - 250px)' }}>
        <iframe
          src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3933.863569644036!2d103.7715890900215!3d1.3061327858715956!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x31da1af50a2f1ebf%3A0x8aea55fe34ee4a51!2sTembusu%20College!5e0!3m2!1sen!2ssg!4v1593365665107!5m2!1sen!2ssg"
          width="100%"
          height="100%"
          frameBorder="0"
          style={{ border: '0' }}
          allowFullScreen={false}
          aria-hidden="false"
          tabIndex={0}
        ></iframe>
      </div>
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
