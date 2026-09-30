<script lang="ts">
  import Router from "svelte-spa-router";
  import Home from './pages/Home.svelte';
  import Profile from './pages/Profile.svelte';

  const routes = {
    '/' : Home,
    '/profile' : Profile,
  }
</script>

<Router routes={routes}/>