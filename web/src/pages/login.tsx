import React, { useState } from 'react';
import type { NextPage } from 'next';
import Link from 'next/link';
import MobileDetect from 'mobile-detect';
import {
  Button,
  Container,
  Divider,
  Grid,
  Header,
  Icon,
  Image,
  List,
  Menu,
  Responsive,
  Segment,
  Sidebar,
  Visibility,
  GridColumn,
  Input,
} from 'semantic-ui-react';

const getWidthFactory = (isMobileFromSSR) => (): number => {
  const isSSR = typeof window === 'undefined';
  const ssrValue: number = isMobileFromSSR
    ? (Responsive.onlyMobile.maxWidth as number)
    : (Responsive.onlyTablet.minWidth as number);
  return isSSR ? ssrValue : window.innerWidth;
};

type HomepageHeadingProps = {
  mobile?: boolean;
};

const HomepageHeading = ({ mobile }: HomepageHeadingProps) => (
  <Container text>
    <Header
      as="h1"
      content="Rankathon"
      inverted
      style={{
        fontSize: mobile ? '2em' : '4em',
        fontWeight: 'normal',
        marginBottom: 0,
        marginTop: mobile ? '1.5em' : '3em',
      }}
    />
    <Header
      as="h2"
      content="Smart, algorithm powered hackathons."
      inverted
      style={{
        fontSize: mobile ? '1.5em' : '1.7em',
        fontWeight: 'normal',
        marginTop: mobile ? '0.5em' : '1.5em',
      }}
    />
    <Button primary size="huge">
      Get Started
      <Icon name="arrow right" />
    </Button>
  </Container>
);

/* Heads up!
 * Neither Semantic UI nor Semantic UI React offer a responsive navbar, however, it can be implemented easily.
 * It can be more complicated, but you can create really flexible markup.
 */

function DesktopContainer({ getWidth, children }) {
  const [fixed, setFixed] = useState<boolean>(false);

  function hideFixedMenu() {
    setFixed(false);
  }

  function showFixedMenu() {
    setFixed(true);
  }

  return (
    <Responsive getWidth={getWidth} minWidth={Responsive.onlyTablet.minWidth}>
      <Visibility
        once={false}
        onBottomPassed={showFixedMenu}
        onBottomPassedReverse={hideFixedMenu}
      >
        <Segment
          inverted
          textAlign="center"
          style={{ minHeight: 700, padding: '1em 0em' }}
          vertical
        >
          <Menu
            fixed={fixed ? 'top' : null}
            inverted={!fixed}
            pointing={!fixed}
            secondary={!fixed}
            size="large"
          >
            <Container>
              <Menu.Item as="a" active>
                Home
              </Menu.Item>
              <Menu.Item as="a" href="about">
                About
              </Menu.Item>
              <Menu.Item as="a">Try</Menu.Item>
              <Menu.Item as="a">Contact</Menu.Item>
              <Menu.Item position="right">
                <Button as="a" inverted={!fixed}>
                  Log in
                </Button>
                <Button
                  as="a"
                  inverted={!fixed}
                  primary={fixed}
                  style={{ marginLeft: '0.5em' }}
                >
                  Sign Up
                </Button>
              </Menu.Item>
            </Container>
          </Menu>
          <HomepageHeading />
        </Segment>
      </Visibility>

      {children}
    </Responsive>
  );
}

type MobileContainerProps = {
  getWidth: () => number;
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

type ResponsiveContainerProps = {
  getWidth: () => number;
};

const ResponsiveContainer: React.FC<ResponsiveContainerProps> = ({
  getWidth,
  children,
}) => (
  <div>
    <DesktopContainer getWidth={getWidth}>{children}</DesktopContainer>
    <MobileContainer getWidth={getWidth}>{children}</MobileContainer>
  </div>
);

type PageProps = {
  getWidth: () => number;
};

const HomepageLayout: NextPage<PageProps> = ({ getWidth }) => (
  <ResponsiveContainer getWidth={getWidth}>
    <Segment vertical style={{ marginTop: '2em' }}>
      <Grid textAlign="center" verticalAlign="middle">
        <Grid.Row>
          <Grid.Column>
            <p>
              <Image
                rounded
                size="small"
                src="img/icon.svg"
                style={{ margin: '0.7em auto' }}
              />
            </p>
            <p>
              <Input
                size="large"
                style={{ margin: '0.5em auto', width: '70%' }}
                placeholder="Username"
              />
            </p>
            <p>
              <Input
                size="large"
                style={{ margin: '0.5em auto', width: '70%' }}
                placeholder="Password"
              />
            </p>
            <p>
              <Button primary style={{ margin: '1em 0.5em' }} size="huge">
                Login
              </Button>
              <Button secondary style={{ margin: '1em 0.5em' }} size="huge">
                Sign-Up
              </Button>
            </p>
            <p>Forgot Password?</p>
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </Segment>
  </ResponsiveContainer>
);

HomepageLayout.getInitialProps = async ({ req }) => {
  const result: MobileDetect = new MobileDetect(req.headers['user-agent']);
  const isMobile: boolean = !!result.mobile();
  return { getWidth: getWidthFactory(isMobile) };
};

export default HomepageLayout;
