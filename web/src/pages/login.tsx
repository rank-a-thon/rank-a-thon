import React, { useState } from 'react';
import type { NextPage } from 'next';
import Link from 'next/link';
import {
  Button,
  Grid,
  Icon,
  Image,
  Menu,
  Responsive,
  Segment,
  Sidebar,
  GridColumn,
  Input,
} from 'semantic-ui-react';

type MobileContainerProps = {
  getWidth?: () => number;
  children: React.ReactNode;
};

function MobileContainer(props: MobileContainerProps) {
  const [sidebarOpened, setSidebarOpened] = useState<boolean>(false);
  function handleSidebarHide() {
    setSidebarOpened(false);
  }
  function handleToggle() {
    setSidebarOpened(true);
  }

  return (
    <Responsive
      as={Sidebar.Pushable}
      getWidth={props.getWidth}
      maxWidth={Responsive.onlyMobile.maxWidth}
    >
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
          style={{ minHeight: 70, padding: '2em 0em' }}
          vertical
        >
          <Grid columns="equal">
            <Grid.Row>
              <GridColumn>
                <Icon onClick={handleToggle} name="angle double left" />
              </GridColumn>
              <GridColumn as="h3">Login</GridColumn>
              <GridColumn></GridColumn>
            </Grid.Row>
          </Grid>
        </Segment>
        {props.children}
      </Sidebar.Pusher>
    </Responsive>
  );
}

type PageProps = {
  getWidth?: () => number;
};

const LoginLayout: NextPage<PageProps> = () => {
  const [authAction, setAuthAction] = useState<'login' | 'signup'>('login');

  function handleTabClick(e, { name }) {
    setAuthAction(name);
  }

  return (
    <MobileContainer>
      <Segment vertical style={{ marginTop: '2em' }}>
        <Grid textAlign="center" verticalAlign="middle">
          <Grid.Row>
            <Grid.Column>
              <p>
                <Image
                  rounded
                  size="small"
                  src="img/icon.svg"
                  style={{ margin: '0 auto' }}
                />
              </p>
              <Menu tabular widths={2}>
                <Menu.Item
                  name="login"
                  active={authAction === 'login'}
                  onClick={handleTabClick}
                />
                <Menu.Item
                  name="signup"
                  active={authAction === 'signup'}
                  onClick={handleTabClick}
                />
              </Menu>

              {authAction === 'login' && (
                <>
                  <p>
                    <Input
                      size="large"
                      style={{ margin: '0.5em auto', width: '70%' }}
                      placeholder="Email"
                    />
                  </p>
                  <p>
                    <Input
                      size="large"
                      style={{ margin: '0.5em auto', width: '70%' }}
                      placeholder="Password"
                    />
                  </p>
                  <Button primary style={{ margin: '1em 0.5em' }} size="huge">
                    Login
                  </Button>
                  <p>Forgot Password?</p>
                </>
              )}

              {authAction === 'signup' && (
                <>
                  <p>
                    <Input
                      size="large"
                      style={{ margin: '0.2em auto', width: '70%' }}
                      placeholder="Name"
                    />
                  </p>
                  <p>
                    <Input
                      size="large"
                      style={{ margin: '0.2em auto', width: '70%' }}
                      placeholder="Email"
                    />
                  </p>
                  <p>
                    <Input
                      size="large"
                      style={{ margin: '0.2em auto', width: '70%' }}
                      placeholder="Password"
                    />
                  </p>
                  <p>
                    <Input
                      size="large"
                      style={{ margin: '0.2em auto', width: '70%' }}
                      placeholder="Repeat Password"
                    />
                  </p>
                  <p>
                    <Button
                      secondary
                      style={{ margin: '1em 0.5em' }}
                      size="huge"
                    >
                      Sign-Up
                    </Button>
                  </p>
                </>
              )}
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Segment>
    </MobileContainer>
  );
};

export default LoginLayout;
