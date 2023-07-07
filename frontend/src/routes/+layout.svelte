<script lang="ts">
    import "../app.css";

    let sidebarWidth = 300;
    let sidebarOpened = true;
    let sidebarMoving = false;
    let toolbarWidth: number;
    let contentWidth: number;
    let contentTag: HTMLElement;

    function toggleSidebar() {
        sidebarOpened = !sidebarOpened;
        sidebarMoving = true;
    }

    $: if (!sidebarOpened) {
        toolbarWidth = 100;
        contentWidth = 0;
    } else {
        toolbarWidth = sidebarWidth;
        contentWidth = sidebarWidth;
    }

    function onAnimationEnd() {
        contentTag.classList.remove('transition');
        sidebarMoving = false;
    }
</script>

<main class="h-screen w-screen relative">
    <div class="fixed h-[38px] w-screen z-50 flex items-center bg-transparent" style="--wails-draggable:drag">
        <div id="toolbar" class:closing={sidebarMoving && !sidebarOpened} class:oping={sidebarMoving && sidebarOpened} class="pl-[77px] flex-none" style="width: {toolbarWidth}px;">
            <button on:click={toggleSidebar}>T</button>
        </div>
        <div class="flex justify-between w-full">
            <div>prefix</div>
            <div>suffix</div>
        </div>
    </div>
    <div class="h-full flex flex-col z-0" style="width: {sidebarWidth}px;">
        <div class="h-[38px] flex-none"></div>
        <div class="h-full flex justify-between items-center">
            <p>Sidebar Left</p>
            <span>Sidebar Right</span>
            <div class="w-10 h-10 bg-blue-200"></div>
        </div>
    </div>
    <div id="content" bind:this={contentTag} class="absolute z-40 mt-[-100vh] float-right right-0 bg-white h-full flex flex-col"
         class:expanded={!sidebarOpened}
         class:transition={sidebarMoving}
         on:transitionend={onAnimationEnd}
         style="width: calc(100vw - {contentWidth}px);">
        <div class="h-[38px] flex-none"></div>
        {sidebarMoving}
        <slot/>
        <button class="border" on:click={toggleSidebar}>T</button>
    </div>
</main>

<style lang="postcss">
    #content.expanded {
        width: 100vw;
    }

    #content.transition {
        transition: width 400ms ease-in-out;
        box-shadow: 2px 0 50px rgba(0, 0, 0, 0.25);
    }

    #toolbar.closing {
        transition: width 350ms ease-in-out;
    }

    #toolbar.oping {
        transition: width 400ms ease-in-out;
    }
</style>
