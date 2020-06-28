import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import {
  Grid,
  Icon,
  Menu,
  Segment,
  Sidebar,
  GridColumn,
} from 'semantic-ui-react';

type MobileContainerProps = {
  getWidth?: () => number;
  children: React.ReactNode;
  title: string;
};

function MobilePostAuthContainer(props: MobileContainerProps) {
  const [sidebarOpened, setSidebarOpened] = useState<boolean>(false);
  function handleSidebarHide() {
    setSidebarOpened(false);
  }
  function handleToggle() {
    setSidebarOpened(true);
  }

  return (
    <>
      <Sidebar.Pushable style={{ height: '100vh' }}>
        <Sidebar
          as={Menu}
          animation="push"
          inverted
          onHide={handleSidebarHide}
          vertical
          visible={sidebarOpened}
        >
          <Menu.Item active>
            <Link href="/">
              <a>Home</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="">
              <a>About</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="">
              <a>Try</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="">
              <a>Contact</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="">
              <a>Log in</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="">
              <a>Sign Up</a>
            </Link>
          </Menu.Item>
        </Sidebar>

        <Sidebar.Pusher dimmed={sidebarOpened}>
          <Segment
            inverted
            textAlign="center"
            style={{
              minHeight: 70,
              padding: '2em 0em',
            }}
            vertical
          >
            <Grid columns="equal">
              <Grid.Row>
                <GridColumn>
                  <Icon onClick={handleToggle} name="angle double left" />
                </GridColumn>
                <GridColumn as="h3">{props.title}</GridColumn>
                <GridColumn></GridColumn>
              </Grid.Row>
            </Grid>
          </Segment>
          <div style={{ height: '100%', marginBottom: '63px' }}>
            {props.children}
          </div>
        </Sidebar.Pusher>
      </Sidebar.Pushable>
      <Menu widths={4} icon="labeled" fixed="bottom" size="tiny">
        <Menu.Item name="Explore">
          <Icon name="search" />
          Explore
        </Menu.Item>
        <Menu.Item name="Schedule">
          <Icon name="time" />
          Schedule
        </Menu.Item>
        <Menu.Item name="Map">
          <Icon name="map" />
          Map
        </Menu.Item>
        <Menu.Item name="Team">
          <Icon name="group" />
          Team
        </Menu.Item>
      </Menu>
    </>
  );
}

export default MobilePostAuthContainer;
