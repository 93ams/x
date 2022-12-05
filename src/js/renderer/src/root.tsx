import { component$, $ } from '@builder.io/qwik';
import { QwikCityProvider, RouterOutlet, ServiceWorkerRegister } from '@builder.io/qwik-city';
import { RouterHead } from './components/router-head/router-head';
import { UrqlProvider } from 'qwik-urql';
import {clientFactory} from "~/client";
import '~/styles';

export default component$(() => {
  return <UrqlProvider client={$(clientFactory)}>
      <QwikCityProvider>
        <head>
          <meta charSet="utf-8" />
          <link rel="manifest" href="/src/js/renderer/public/manifest.json" />
          <RouterHead />
        </head>
        <body lang="en">
          <RouterOutlet />
          <ServiceWorkerRegister />
        </body>
      </QwikCityProvider>
    </UrqlProvider>
})
