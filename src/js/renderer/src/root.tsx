import { component$, useStyles$ } from '@builder.io/qwik';
import { QwikCityProvider, RouterOutlet, ServiceWorkerRegister } from '@builder.io/qwik-city';
import { RouterHead } from './components/router-head/router-head';
import globalStyles from './global.css?inline';
import { UrqlProvider } from 'qwik-urql';
import {clientFactory} from "~/client";

export default component$(() => {
  useStyles$(globalStyles);

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
