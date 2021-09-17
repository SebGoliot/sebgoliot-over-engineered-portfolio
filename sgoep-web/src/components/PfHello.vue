<template>

  <section>
    <header>
      <span class="material-icons-round">face</span>
      <h2>&nbsp;Hello !</h2>
    </header>
    <p class="section-margin">
      {{about}}
    </p>
  </section>

</template>

<script>
export default {
  name: "PfHello",
  props: {
    about: {},
  }
};
</script>
