<template>
  <img class="wave-to-blue" src="../assets/svg/wave-to-blue.svg" alt="" />

  <footer>
    <div>
      <p>
        Most icons are from
        <a href="https://icons8.com/" target="_blank" rel="noopener noreferrer"
          >Icons8</a
        >
      </p>
      <p>
        The other ones are from
        <a
          href="https://fonts.google.com/icons"
          target="_blank"
          rel="noopener noreferrer"
          >Google Fonts</a
        >
      </p>
    </div>
  </footer>
</template>

<script>
export default { name: "Footer" };
</script>

<style scoped>
footer {
  background-color: var(--primary);
  color: var(--light);
}
footer > div {
  margin: auto;
  max-width: 900px;
  display: flex;
  flex-flow: row wrap;
  justify-content: space-evenly;
}
</style>
