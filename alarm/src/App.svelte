<script>
  import './css/global.scss';
  import { mdiAlarm } from '@mdi/js';
  import Icon from './lib/Icon.svelte';
  import { onMount } from 'svelte';

  let className;
  let timing;

  onMount(() => {
    const url = new URL(window.location.href);
    const params = new URLSearchParams(url.search);

    className = params.get('className');
    timing = params.get('timing');
  });
</script>

<main>
  <br /><br /><br /><br />
  <div class="icon">
    <Icon size={90} path={mdiAlarm} />
  </div>
  <br /><br />
  <h1>Your class <mark>{className}</mark> is scheduled @ <mark>{timing}</mark></h1>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .icon {
    position: relative;

    width: 7rem;
    height: 7rem;

    display: flex;
    justify-content: center;
    align-items: center;
  }

  .icon::before {
    content: '';

    position: absolute;
    top: 0;
    left: 0;
    z-index: -1;

    border-radius: 50%;

    width: inherit;
    height: inherit;

    background-color: rgba(var(--app-color-dark-rgb), 0.4);
    animation: ripple 1.5s infinite;
    animation-timing-function: ease;
  }

  .icon > :global(svg) {
    fill: var(--app-color-dark);
    animation: shake 2.5s infinite;
    animation-delay: 100ms;
  }

  h1 {
    font-size: 4rem;
    font-weight: 600;

    line-height: 1.4;

    text-align: center;

    width: min(60rem, 100vw);
  }

  h1 mark {
    background-color: transparent;
    color: var(--app-color-primary);
  }

  @keyframes shake {
    10%,
    90% {
      transform: translate3d(-1px, 0, 0);
    }

    20%,
    80% {
      transform: translate3d(2px, 0, 0);
    }

    30%,
    50%,
    70% {
      transform: translate3d(-4px, 0, 0);
    }

    40%,
    60% {
      transform: translate3d(4px, 0, 0);
    }
  }

  @keyframes ripple {
    0% {
      transform: scale(1);
      opacity: 1;
    }

    60% {
      transform: scale(2, 2);
      opacity: 0;
    }

    100% {
      transform: scale(2, 2);
      opacity: 0;
    }
  }

  @media screen and (max-width: 600px) {
    h1 {
      font-size: 2.618rem;
    }
  }
</style>
