<script setup lang="ts">

</script>

<template>
   <div class="fixed left-1/2 -translate-x-1/2 bottom-44 z-49 bg-zinc-800 text-zinc-100 border border-gray-900/50 rounded-md px-4 py-2 shadow-lg flex items-start space-x-2 text-sm animate-slide-in">
        <span class="mr-0.5">
            ⚠️
        </span>
        <span>
            Beberapa model AI <span class="font-semibold text-yellow-400">lambat</span> merespons.<br>
            Gunakan <span class="font-bold text-indigo-400">Gemini</span> 🚀 untuk hasil terbaik.
        </span>
        <button class="ml-2 text-zinc-400 hover:text-white" onclick="this.parentElement.remove()">
            <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd"
                    d="M10 8.586l4.95-4.95 1.414 1.414L11.414 10l4.95 4.95-1.414 1.414L10 11.414l-4.95 4.95-1.414-1.414L8.586 10 3.636 5.05l1.414-1.414L10 8.586z"
                    clip-rule="evenodd" />
                </svg>
        </button>
    </div>
</template>