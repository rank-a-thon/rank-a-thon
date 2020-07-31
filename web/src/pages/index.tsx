import React, { useState, useEffect } from 'react';
import type { NextPage } from 'next';
import Link from 'next/link';
import { useRouter } from 'next/router';
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
    <Link href="/signup">
      <Button as="a" primary size="huge">
        Sign Up Now
        <Icon name="arrow right" />
      </Button>
    </Link>
  </Container>
);

/* Heads up!
 * Neither Semantic UI nor Semantic UI React offer a responsive navbar, however, it can be implemented easily.
 * It can be more complicated, but you can create really flexible markup.
 */

function useMobilePlease() {
  if (
    window !== undefined &&
    window.innerWidth >= Responsive.onlyMobile.maxWidth
  ) {
    alert(
      'Sorry, but this app is only available in mobile for now! Please view this app on your phone OR use your browser developer tools to set a mobile device emulation!',
    );
  }
}
function DesktopContainer({ getWidth, children }) {
  const [fixed, setFixed] = useState<boolean>(false);

  function hideFixedMenu() {
    setFixed(false);
  }

  function showFixedMenu() {
    setFixed(true);
  }

  useEffect(useMobilePlease, []);

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
            <Container onClick={useMobilePlease}>
              <Menu.Item as="a" active>
                Home
              </Menu.Item>
              <Menu.Item as="a">About</Menu.Item>
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
        <SidebarItem name="Home" href="/" />
        <SidebarItem name="Log in" href="/login" />
        <SidebarItem name="Sign Up" href="/signup" />
      </Sidebar>

      <Sidebar.Pusher dimmed={sidebarOpened}>
        <Segment
          inverted
          textAlign="center"
          style={{ minHeight: 350, padding: '1em 0em' }}
          vertical
        >
          <Container>
            <Menu inverted pointing secondary size="large">
              <Menu.Item onClick={handleToggle}>
                <Icon name="sidebar" />
              </Menu.Item>
              <Menu.Item position="right">
                <Link href="login">
                  <Button as="a" inverted>
                    Log in
                  </Button>
                </Link>
                <Link href="/signup">
                  <Button as="a" inverted style={{ marginLeft: '0.5em' }}>
                    Sign Up
                  </Button>
                </Link>
              </Menu.Item>
            </Menu>
          </Container>
          <HomepageHeading mobile />
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
    <Segment style={{ padding: '4em 0em' }} vertical>
      <Grid container stackable verticalAlign="middle">
        <Grid.Row>
          <Grid.Column width={8}>
            <Header as="h3" style={{ fontSize: '2em' }}>
              Smart, Reliable, Efficient Hackathons
            </Header>
            <p style={{ fontSize: '1.33em' }}>
              Rank-a-thon is E2E, B2B web application that manages hackathons.
              It allows for algorithmic judging based on pairwise ranking
              algorithms. This reduces subjectivity, enhances reliability,
              efficiency, and transparency of the judging process.
            </p>
            <Header as="h3" style={{ fontSize: '2em' }}>
              Features
            </Header>
            <p style={{ fontSize: '1.33em' }}>
              <ul>
                <li> Find hackathon key information </li>
                <li> View other team’s projects </li>
                <li> Easily find teams assigned for judging </li>
                <li> Judge teams based on rankings </li>
                <li> Live-generate prize winners </li>
              </ul>
            </p>
          </Grid.Column>
          <Grid.Column floated="right" width={6}>
            <Image rounded size="large" src="img/logo.svg" />
          </Grid.Column>
        </Grid.Row>
        <Grid.Row>
          <Grid.Column textAlign="center">
            <Button size="huge">Sign Up Now</Button>
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </Segment>

    <Segment style={{ padding: '0em' }} vertical>
      <Grid celled="internally" columns="equal" stackable>
        <Grid.Row textAlign="center">
          <Grid.Column style={{ paddingBottom: '5em', paddingTop: '5em' }}>
            <Header as="h3" style={{ fontSize: '2em' }}>
              "What a Company"
            </Header>
            <p style={{ fontSize: '1.33em' }}>
              That is what they all say about us
            </p>
          </Grid.Column>
          <Grid.Column style={{ paddingBottom: '5em', paddingTop: '5em' }}>
            <Header as="h3" style={{ fontSize: '2em' }}>
              "I shouldn't have gone with their competitor."
            </Header>
            <p style={{ fontSize: '1.33em' }}>
              <Image avatar src="img/icon.svg" />
              <b>Nan</b> Chief Fun Officer Acme Toys
            </p>
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </Segment>

    <Segment style={{ padding: '8em 0em' }} vertical>
      <Container text>
        <Header as="h3" style={{ fontSize: '2em' }}>
          Motivation
        </Header>
        <p style={{ fontSize: '1.33em' }}>
          In large-scale hackathons such as Hack&amp;Roll (a hackathon organised
          by NUS Hackers which has over 500 participants), it is always a
          challenge to ensure that teams are judged sufficiently, fairly, and
          quickly. Teams need to be given enough time to talk about their work
          and showcase their product, and need to be given more opportunities to
          present their work to judges of different backgrounds, to reduce bias
          and variance in the judging process. However, judging at such
          hackathons often have tight time and manpower constraints, which
          greatly increases the complexity of such a large-scale judging
          operation.
        </p>
        <Button as="a" size="large">
          Read More
        </Button>

        <Divider
          as="h4"
          className="header"
          horizontal
          style={{ margin: '3em 0em', textTransform: 'uppercase' }}
        >
          <Link href="">
            <a href="#">Case Studies</a>
          </Link>
        </Divider>

        <Header as="h3" style={{ fontSize: '2em' }}>
          Hack&amp;Roll 2021
        </Header>
        <p style={{ fontSize: '1.33em' }}>
          To be deployed for Hack&amp;Roll 2021, Singapore’s largest student-run
          hackathon.
        </p>
        <Button as="a" size="large">
          Read More
        </Button>
      </Container>
    </Segment>

    <Segment inverted vertical style={{ padding: '5em 0em' }}>
      <Container>
        <Grid divided inverted stackable>
          <Grid.Row>
            <Grid.Column width={3}>
              <Header inverted as="h4" content="About" />
              <List link inverted>
                <List.Item as="a">About</List.Item>
                <List.Item as="a">Contact Us</List.Item>
                <List.Item as="a">Sitemap</List.Item>
              </List>
            </Grid.Column>
            <Grid.Column width={7}>
              <Header as="h4" inverted>
                &copy; Rankathon 2020
              </Header>
              <p>We Rank Hackathons</p>
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Container>
    </Segment>
  </ResponsiveContainer>
);

HomepageLayout.getInitialProps = async ({ req }) => {
  if (!req) {
    return { getWidth: getWidthFactory(true) };
  }
  const result: MobileDetect = new MobileDetect(req.headers['user-agent']);
  const isMobile: boolean = !!result.mobile();
  return { getWidth: getWidthFactory(isMobile) };
};

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

export default HomepageLayout;
