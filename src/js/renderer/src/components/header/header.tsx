import { component$ } from '@builder.io/qwik';

export default component$(() => {
  return (
    <header>
      <div class="logo">
        <a href="/">
          Home
        </a>
      </div>
      <ul>
        <li>
          <a href="/docs" class='btn'>
            Docs
          </a>
        </li>
        <li>
          <a href="/flower" class='btn'>
            Examples
          </a>
        </li>
        <li>
          <a href="/tutorials" class='btn'>
            Tutorials
          </a>
        </li>
      </ul>
    </header>
  );
});
