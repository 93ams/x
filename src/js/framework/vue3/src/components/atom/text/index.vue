<template>
  <div>
    {{ message || "Hello" }} {{ name }}! I can run in React, Vue, Solid or
    Svelte!
  </div>
</template>

<script lang="ts">
type Props = {
  message: string;
};

export default {
  name: "text",

  props: ["message"],

  data: () => ({ name: "Foo" }),
};
</script>