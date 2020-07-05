import React, { useState, useEffect } from 'react';
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

import { clearMe, getMe } from '../data/me';

type MobileContainerProps = {
  getWidth?: () => number;
  children: React.ReactNode;
  title: string;
  requireAuth?: boolean;
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

  useEffect(() => {
    // Hook to check for auth
    // If user is not authed, kick them back to login screen
    if (props.requireAuth && !getMe()) {
      logOut();
    }
  }, []);

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
          <SidebarItem name="Dashboard" href="/dashboard" />
          <SidebarItem name="Announcements" href="/announcements" />
          <SidebarItem name="Manage Team" href="/team" />
          <SidebarItem name="View Projects" href="/explore" />
          <SidebarItem name="Map" href="/map" />
          <SidebarItem name="Schedule" href="/schedule" />
          <SidebarItem name="Log Out" onClick={logOut} />
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
                  <Icon
                    onClick={handleToggle}
                    style={{
                      padding: '1em',
                      margin: '-1em',
                    }}
                    name="sidebar"
                  />
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

type SidebarItemProps = {
  href?: string;
  name: string;
  onClick?: any;
};

function SidebarItem(props: SidebarItemProps): JSX.Element {
  const router = useRouter();
  const isActive = router.pathname === props.href;
  if (props.href && props.onClick) {
    throw "Can't have both href and onClick";
  }
  if (props.onClick) {
    return <Menu.Item onClick={props.onClick}>{props.name}</Menu.Item>;
  }
  return (
    <Link href={props.href}>
      <Menu.Item active={isActive}>{props.name}</Menu.Item>
    </Link>
  );
}

export default MobilePostAuthContainer;
