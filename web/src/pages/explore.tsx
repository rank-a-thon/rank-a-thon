import React from 'react';
import type { NextPage } from 'next';
import { Segment, Card } from 'semantic-ui-react';
import MobilePostAuthContainer from '../components/MobilePostAuthContainer';

type PageProps = {
  getWidth?: () => number;
};

const DashboardLayout: NextPage<PageProps> = () => {
  return (
    <MobilePostAuthContainer title="Explore">
      <Segment
        basic
        textAlign="left"
        style={{ padding: '1.5em 2em 0.8em 2em' }}
      >
        <Card
          image="/img/project_1.webp"
          header="AutoModReg"
          meta="by 2GUD4SKOOL"
          description="Tool that chooses your modules at random. Our tool uses quantum numpy, an next-generation blockchain powered numeric processing library to determine the modules it will pick at random. LUL"
          fluid
        />
        <Card
          image="/img/pepekip.png"
          header="PepeMudKip"
          meta="by Mudkip Lovers"
          description="Difficulty choosing your starter pokemon? PepeMudkip is a new-age full-stack PWA that uses modern natural language processing techniques for text-generation to perform named entity recognition on geographical noise data to help you determine the best starter pokemon to choose."
          fluid
        />
        <Card
          image="/img/hammer.jpeg"
          header="ElectionMaster"
          meta="by MerryGandering"
          description="A handshake 🤝 is basically a 🙏promise, a commitment 💍, a tall order 🍾, means I must meet 🏃‍♂️that tall order , and it's for YOU ☝️! And it's for you, in sense that three👌 fingers also pointing me 🙆‍♂️, it's also for ME! 👇 It's for us. 👨‍👩‍👧‍👧And if the result is good 👍👍👍, THUMBS UP MAAAAN... 👍👍👍 and if the result is lousy 😤😤 wad happen? 🧐𝔁𝓾𝓮🥶𝓱𝓾𝓪🧚‍♀️𝓹𝓲𝓪𝓸😻𝓹𝓲𝓪𝓸🗿𝓫𝓮𝓲👺𝓯𝓮𝓷𝓰🤩𝔁𝓲𝓪𝓸😼𝔁𝓲𝓪𝓸👣"
          fluid
        />
      </Segment>
    </MobilePostAuthContainer>
  );
};

export default DashboardLayout;
