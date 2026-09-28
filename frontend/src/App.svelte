
<script lang="ts">

  let msg = $state('hello world!');
  let toaster: ReturnType<typeof setTimeout>;
  let resultEl: HTMLElement;
  let toastEl: HTMLElement;

  function greet() {
    alert(msg);
  }

  function showToaster(msg: string): void {
    resultEl.innerText = msg;
    toastEl.classList.add('is-visible');
    clearTimeout(toaster);
    toaster = setTimeout(() => {
    toastEl.classList.remove('is-visible');
    }, 3000); 
}
</script>

<!-- svelte-ignore event_directive_deprecated -->
<main>

<button on:click={() => showToaster('Hello, world!')}>Click me</button>
<div class="toast" bind:this={toastEl}>
    <span>from algo play</span>
    <span bind:this={resultEl}></span>
</div>
</main>

<style>
  button {
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
    color: white;
    background-color: #007BFF;
  }
</style>