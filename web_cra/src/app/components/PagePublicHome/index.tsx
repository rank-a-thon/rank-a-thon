/**
 *
 * PagePublicHome
 *
 */
import React, { useState } from 'react';
import styled from 'styled-components/macro';
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
import { Link, NavLink, useHistory } from 'react-router-dom';

interface Props {}

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
    <NavLink to="/signup">
      <Button as="a" primary size="huge">
        Sign Up Now
        <Icon name="arrow right" />
      </Button>
    </NavLink>
  </Container>
);

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
            fixed={fixed ? 'top' : undefined}
            inverted={!fixed}
            pointing={!fixed}
            secondary={!fixed}
            size="large"
          >
            <Container>
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
                <NavLink to="login">
                  <Button as="a" inverted>
                    Log in
                  </Button>
                </NavLink>
                <NavLink to="/signup">
                  <Button as="a" inverted style={{ marginLeft: '0.5em' }}>
                    Sign Up
                  </Button>
                </NavLink>
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

export const PagePublicHome: React.FC = () => (
  <ResponsiveContainer getWidth={() => window.innerWidth}>
    <Segment style={{ padding: '8em 0em' }} vertical>
      <Grid container stackable verticalAlign="middle">
        <Grid.Row>
          <Grid.Column width={8}>
            <Header as="h3" style={{ fontSize: '2em' }}>
              We Help Companies and Companions
            </Header>
            <p style={{ fontSize: '1.33em' }}>
              We can give your company superpowers to do things that they never
              thought possible. Let us delight your customers and empower your
              needs... through pure data analytics.
            </p>
            <Header as="h3" style={{ fontSize: '2em' }}>
              We Make Bananas That Can Dance
            </Header>
            <p style={{ fontSize: '1.33em' }}>
              Yes that's right, you thought it was the stuff of dreams, but even
              bananas can be bioengineered.
            </p>
          </Grid.Column>
          <Grid.Column floated="right" width={6}>
            <Image rounded size="large" src="img/logo.svg" />
          </Grid.Column>
        </Grid.Row>
        <Grid.Row>
          <Grid.Column textAlign="center">
            <Button size="huge">Check Them Out</Button>
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
          Breaking The Grid, Grabs Your Attention
        </Header>
        <p style={{ fontSize: '1.33em' }}>
          Instead of focusing on content creation and hard work, we have learned
          how to master the art of doing nothing by providing massive amounts of
          whitespace and generic content that can seem massive, monolithic and
          worth your attention.
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
          <Link to="">
            <a href="#">Case Studies</a>
          </Link>
        </Divider>

        <Header as="h3" style={{ fontSize: '2em' }}>
          Did We Tell You About Our Bananas?
        </Header>
        <p style={{ fontSize: '1.33em' }}>
          Yes I know you probably disregarded the earlier boasts as non-sequitur
          filler content, but it's really true. It took years of gene splicing
          and combinatory DNA research, but our bananas can really dance.
        </p>
        <Button as="a" size="large">
          I'm Still Quite Interested
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
                <List.Item as="a">Sitemap</List.Item>
                <List.Item as="a">Contact Us</List.Item>
                <List.Item as="a">Religious Ceremonies</List.Item>
                <List.Item as="a">Gazebo Plans</List.Item>
              </List>
            </Grid.Column>
            <Grid.Column width={3}>
              <Header inverted as="h4" content="Services" />
              <List link inverted>
                <List.Item as="a">Banana Pre-Order</List.Item>
                <List.Item as="a">DNA FAQ</List.Item>
                <List.Item as="a">How To Access</List.Item>
                <List.Item as="a">Favorite X-Men</List.Item>
              </List>
            </Grid.Column>
            <Grid.Column width={7}>
              <Header as="h4" inverted>
                &copy; Rankathon 2020
              </Header>
              <p>Orbital is hard.</p>
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Container>
    </Segment>
  </ResponsiveContainer>
);

type SidebarItemProps = {
  href?: string;
  name: string;
  onClick?: any;
};
function SidebarItem(props: SidebarItemProps): JSX.Element {
  const router = useHistory();
  const isActive = window.location.pathname === props.href;
  if (props.href && props.onClick) {
    throw "Can't have both href and onClick";
  }
  if (props.onClick) {
    return <Menu.Item onClick={props.onClick}>{props.name}</Menu.Item>;
  }
  return (
    <NavLink to={props.href as string}>
      <Menu.Item active={isActive}>{props.name}</Menu.Item>
    </NavLink>
  );
}
