<script>
    // import { animate } from 'svelte/animate';

    let isAnimating = false;
    let status = '动画';

    function startAnimation() {
        isAnimating = true;
        status = '动画开始';
    }

    function onAnimationEnd() {
        isAnimating = false;
        status = '动画结束';
    }
</script>

<style>
    .box {
        width: 100px;
        height: 100px;
        background-color: red;
    }

    .animate {
        animation: custom-animation 2s;
    }

    @keyframes custom-animation {
        0% {
            transform: scale(1);
        }
        100% {
            transform: scale(2);
        }
    }
</style>

<!--<div class="box {isAnimating ? 'animate' : ''}"-->
<!--     on:animationstart={startAnimation}-->
<!--     on:animationend={onAnimationEnd}>-->
<!--    <button on:click={() => isAnimating = !isAnimating}>Trigger</button>-->
<!--    <p>{status}</p>-->
<!--</div>-->
