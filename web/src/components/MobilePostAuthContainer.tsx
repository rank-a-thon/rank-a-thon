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

import { clearMe } from '../data/me';

type MobileContainerProps = {
  getWidth?: () => number;
  children: React.ReactNode;
  title: string;
};

function MobilePostAuthContainer(props: MobileContainerProps) {
  const [sidebarOpened, setSidebarOpened] = useState<boolean>(false);
  const router = useRouter();
  function handleSidebarHide() {
    setSidebarOpened(false);
  }
  function handleToggle() {
    setSidebarOpened(true);
  }

  function logOut() {
    clearMe();
    router.push('/login');
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
            <Link href="/dashboard">
              <a>Dashboard</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="/announcements">
              <a>Announcements</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="/team">
              <a>Manage Team</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="/explore">
              <a>View Projects</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="/map">
              <a>Map</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <Link href="/schedule">
              <a>Schedule</a>
            </Link>
          </Menu.Item>
          <Menu.Item>
            <a onClick={logOut}>Log Out</a>
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
        <Link href="/explore">
          <Menu.Item as="a" name="Explore">
            <Icon name="search" />
            Explore
          </Menu.Item>
        </Link>
        <Link href="/schedule">
          <Menu.Item as="a" name="Schedule">
            <Icon name="time" />
            Schedule
          </Menu.Item>
        </Link>
        <Link href="/map">
          <Menu.Item as="a" name="Map">
            <Icon name="map" />
            Map
          </Menu.Item>
        </Link>
        <Link href="/team">
          <Menu.Item as="a" name="Team">
            <Icon name="group" />
            Team
          </Menu.Item>
        </Link>
      </Menu>
    </>
  );
}

export default MobilePostAuthContainer;
