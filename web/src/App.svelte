<script>
  import svelteLogo from "./assets/svelte.svg";
  import Counter from "./lib/Counter.svelte";
  import Todos from "./lib/Todos.svelte";
  import Sample from "./lib/CrudSample.svelte";
  import NewTodo from "./lib/NewTodo.svelte";
  import TodoList from "./lib/TodoList.svelte";
  import EditTodo from "./lib/EditTodo.svelte";

  let editMode = false;

  function handleWelcoming(event) {
		alert(event.detail.message);
	}
</script>

<main>
  <h1>Welcome to Todo (Demo)</h1>
  <!-- <div>
    <a href="https://vitejs.dev" target="_blank" rel="noreferrer"> 
      <img src="/vite.svg" class="logo" alt="Vite Logo" />
    </a>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer"> 
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div>
  <h1>Vite + Svelte</h1> -->

  <div class="card">
    {#if !editMode}
      <NewTodo />
    {:else}
      <EditTodo />
    {/if}
  </div>

  <div class="card">
    <TodoList on:welcoming={handleWelcoming} />
  </div>

  <!-- <div class="card">
    <Todos />
  </div> -->

  <!-- <div class="card">
    <Counter />
  </div> -->

  <!-- <div class="card">
    <Sample />
  </div> -->

  <p>
    Check out <a
      href="https://github.com/sveltejs/kit#readme"
      target="_blank"
      rel="noreferrer">SvelteKit</a
    >, the official Svelte app framework powered by Vite!
  </p>

  <p class="read-the-docs">Click on the Vite and Svelte logos to learn more</p>
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
