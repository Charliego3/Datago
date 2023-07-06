<script lang="ts">
    import "../app.css";

    type SidebarState = 'opened' | 'closed'

    let splitterDragging: boolean = false;
    let currentSidebarWidth: number = 400;
    let sidebarState: SidebarState = 'opened';

    async function toggleSidebar() {
        console.log({
            splitterDragging,
            currentSidebarWidth,
            sidebarState,
        });
    }
</script>

<div class="flex w-screen h-screen overflow-hidden font-robot divide-x divide-gray/35">
    <div class="relative w-[400px] h-full flex-none flex flex-col divide-y divide-gray/35">
        <div class="flex-none h-[38px] flex items-center text-xl">
            <div class="w-[77px] h-full pl-[13px] pr-[11px] flex items-center justify-between">
                <div class="bg-[#FF5F57] w-[12px] h-[12px] rounded-full"/>
                <div class="bg-[#FEBC2E] w-[12px] h-[12px] rounded-full"/>
                <div class="bg-[#27C840] w-[12px] h-[12px] rounded-full"/>
            </div>
            <button class="ml-[80px] fixed z-[999] cursor-default px-[7px] py-[1px] rounded flex hover:bg-slate/15" on:click={toggleSidebar}>
                <span class="i-icons-sidebar-leading opacity-60 cursor-default"/>
                toggle
            </button>
        </div>
        <div class="flex-auto px-[3px]"/>
        <div class="flex-none w-[5px] absolute mt-[-1px] right-[-3px] h-full cursor-col-resize bg-[indigo]/0"></div>
    </div>
    <div id="rightContainer" data-sidebarState={sidebarState} class="h-full flex flex-col divide-y divide-gray/35 z-[9]">
        <div class="flex-none h-[38px] bg-[#FAF9F9] dark:bg-[#373736]" />
        <div class="flex-auto px-[3px] bg-white dark:bg-[#2e2e2d]">
            <slot />
        </div>
    </div>
</div>

<style>
    [data-sidebarState='opened'] {
        width: 100%;
    }

    [data-sidebarState='closed'] {
        width: 100vw;
    }
</style>