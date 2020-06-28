import React from 'react';
import type { NextPage } from 'next';
import { Segment, Item } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  const paragraph = 'Hey';
  return (
    <MobilePostAuthContainer title="Announcements">
      <Item.Group link>
        <Item>
          <Item.Image
            size="tiny"
            src="https://react.semantic-ui.com/images/avatar/large/stevie.jpg"
          />

          <Item.Content>
            <Item.Header>Stevie Feliciano</Item.Header>
            <Item.Description>{paragraph}</Item.Description>
          </Item.Content>
        </Item>

        <Item>
          <Item.Image
            size="tiny"
            src="https://react.semantic-ui.com/images/avatar/large/veronika.jpg"
          />

          <Item.Content>
            <Item.Header>Veronika Ossi</Item.Header>
            <Item.Description>{paragraph}</Item.Description>
          </Item.Content>
        </Item>

        <Item>
          <Item.Image
            size="tiny"
            src="https://react.semantic-ui.com/images/avatar/large/jenny.jpg"
          />

          <Item.Content>
            <Item.Header>Jenny Hess</Item.Header>
            <Item.Description>{paragraph}</Item.Description>
          </Item.Content>
        </Item>
      </Item.Group>
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
