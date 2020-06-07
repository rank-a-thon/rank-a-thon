import 'semantic-ui-css/semantic.min.css';

import { Meta } from '../components/Meta';

function MyApp({ Component, pageProps }) {
  return (
    <>
      <Meta />
      <Component {...pageProps} />
    </>
  );
}

export default MyApp;
