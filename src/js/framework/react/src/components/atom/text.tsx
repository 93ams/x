import * as React from "react";
import { useState } from "react";

type Props = {
  message: string;
};

export default function Text(props: Props) {
  const [name, setName] = useState(() => "Foo");

  return (
    <div>
      {props.message || "Hello"}
      {name}! I can run in React, Vue, Solid or Svelte!
    </div>
  );
}
