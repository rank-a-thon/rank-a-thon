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
  Form,
  Message,
} from 'semantic-ui-react';
import { makeBackendRequest } from '../lib/backend';
import { saveMe } from '../data/me';

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
    <Sidebar.Pushable>
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
    </Sidebar.Pushable>
  );
}

type PageProps = {
  getWidth?: () => number;
};

const LoginLayout: NextPage<PageProps> = () => {
  const [authAction, setAuthAction] = useState<'login' | 'signup'>('login');

  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [passwordCheck, setPasswordCheck] = useState<string>('');
  const [name, setName] = useState<string>('');

  const [error, setError] = useState<{ message: string } | null>(null);
  const [success, setSuccess] = useState<{ message: string } | null>(null);

  function handleTabClick(e, { name }) {
    setError(null);
    setSuccess(null);
    setAuthAction(name);
  }

  function signUp() {
    setError(null);
    setSuccess(null);

    if (!name || !email || !password || !passwordCheck) {
      setError({ message: 'Please fill all fields' });
      return;
    }

    if (password !== passwordCheck) {
      setError({ message: 'Passwords do not match' });
      return;
    }

    if (!email.includes('@')) {
      // TODO: more robust email detection pls
      setError({ message: 'Invalid email' });
      return;
    }

    makeBackendRequest('post', 'v1/user/register', {
      name: name,
      email: email,
      password: password,
    })
      .then(() => {
        setAuthAction('login');
        setSuccess({ message: 'Successfully registered, please login.' });
      })
      .catch((err) => {
        if (err.response && err.response.status === 406) {
          setError({ message: 'Account already exists' });
        } else {
          setError({ message: error.message });
        }
      });
  }

  function login() {
    setError(null);
    setSuccess(null);

    if (!email || !password) {
      setError({ message: 'Please fill all fields' });
      return;
    }

    if (!email.includes('@')) {
      // TODO: more robust email detection pls
      setError({ message: 'Invalid email' });
      return;
    }

    makeBackendRequest('post', 'v1/user/login', {
      email: email,
      password: password,
    })
      .then((response) => {
        setSuccess({ message: 'Successfully logged in. Redirecting.' });
        saveMe(response.data.token, response.data.user);
      })
      .catch((err) => {
        if (err.response) {
          setError({ message: 'Wrong credentials. Please try again.' });
        } else {
          setError({ message: error.message });
        }
      });
  }

  return (
    <MobileContainer>
      <Segment
        basic
        textAlign="center"
        style={{ margin: '2em 0', padding: '0' }}
      >
        <div>
          <Image
            rounded
            size="small"
            src="img/icon.svg"
            style={{ margin: '0 auto' }}
          />
        </div>
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
          <Form>
            <Form.Field>
              <Input
                size="large"
                style={{ margin: '0.5em auto', width: '70%' }}
                placeholder="Email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </Form.Field>
            <Form.Field>
              <Input
                size="large"
                style={{ margin: '0.5em auto', width: '70%' }}
                placeholder="Password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </Form.Field>
            {error && (
              <Message negative style={{ width: '70%', margin: '0 auto' }}>
                {error.message}
              </Message>
            )}
            {success && (
              <Message positive style={{ width: '70%', margin: '0 auto' }}>
                {success.message}
              </Message>
            )}
            <Button
              primary
              style={{ margin: '1em 0.5em' }}
              size="huge"
              onClick={login}
            >
              Login
            </Button>
            <p onClick={() => alert('Not implemented yet')}>Forgot Password?</p>
          </Form>
        )}

        {authAction === 'signup' && (
          <Form>
            <Form.Field>
              <Input
                size="large"
                style={{ margin: '0.2em auto', width: '70%' }}
                placeholder="Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </Form.Field>
            <Form.Field>
              <Input
                size="large"
                style={{ margin: '0.2em auto', width: '70%' }}
                placeholder="Email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </Form.Field>
            <Form.Field>
              <Input
                size="large"
                style={{ margin: '0.2em auto', width: '70%' }}
                placeholder="Password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </Form.Field>
            <Form.Field>
              <Input
                size="large"
                style={{ margin: '0.2em auto', width: '70%' }}
                placeholder="Repeat Password"
                type="password"
                value={passwordCheck}
                onChange={(e) => setPasswordCheck(e.target.value)}
              />
            </Form.Field>
            {error && (
              <Message negative style={{ width: '70%', margin: '0 auto' }}>
                {error.message}
              </Message>
            )}
            <Form.Field>
              <Button
                secondary
                style={{ margin: '1em 0.5em' }}
                size="huge"
                onClick={signUp}
              >
                Sign-Up
              </Button>
            </Form.Field>
          </Form>
        )}
      </Segment>
    </MobileContainer>
  );
};

export default LoginLayout;
