import { component$, useStylesScoped$ } from '@builder.io/qwik';
import { QwikLogo } from '../icons/qwik';
import styles from './header.css?inline';

export default component$(() => {
  useStylesScoped$(styles);

  return (
    <header>
      <div class="logo">
        <a href="~/components/header/header" target="_blank">
          <QwikLogo />
        </a>
      </div>
      <ul>
        <li>
          <a href="~/components/header/header" target="_blank">
            Docs
          </a>
        </li>
        <li>
          <a href="~/components/header/header" target="_blank">
            Examples
          </a>
        </li>
        <li>
          <a href="~/components/header/header" target="_blank">
            Tutorials
          </a>
        </li>
      </ul>
    </header>
  );
});
