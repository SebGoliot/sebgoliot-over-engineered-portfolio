<template>

  <section>
    <header>
      <span class="material-icons-round">face</span>
      <h2>&nbsp;Hello !</h2>
    </header>
    <p class="section-margin">
      <span v-for="(line, key) in about.split('\\n')" :key="key">
        {{ line }} <br />
      </span>
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
