import Head from 'next/head';
import { seoData } from '../data/Seo';

export function Meta(props) {
  return (
    <Head>
      <title>Rankathon</title>
      <meta
        key="description"
        name="description"
        content={seoData.description}
      />
      <meta key="og:type" name="og:type" content={seoData.openGraph.type} />
      <meta key="og:title" name="og:title" content={seoData.openGraph.title} />
      <meta
        key="og:description"
        name="og:description"
        content={seoData.openGraph.description}
      />
      <meta key="og:url" name="og:url" content={seoData.openGraph.url} />
      <meta key="og:image" name="og:image" content={seoData.openGraph.image} />
      <meta
        name="viewport"
        content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no"
      />
    </Head>
  );
}
