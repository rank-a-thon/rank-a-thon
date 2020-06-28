import React from 'react';
import type { NextPage } from 'next';
import { Segment, Item } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';
import announcements from '../data/announcements.json';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  const paragraph = 'Hey';
  return (
    <MobilePostAuthContainer title="Announcements" requireAuth>
      <Item.Group divided>
        {announcements.map((announcement, index) => (
          <Item key={index} style={{ padding: '0' }}>
            <Item.Content
              style={{
                padding: '1.5em',
                backgroundColor:
                  announcement.level === 'urgent' ? 'pink' : null,
              }}
            >
              <Item.Header>{announcement.subject}</Item.Header>
              <Item.Meta>
                {announcement.time} by {announcement.author}
              </Item.Meta>
              <Item.Description>{announcement.body}</Item.Description>
            </Item.Content>
          </Item>
        ))}
      </Item.Group>
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
