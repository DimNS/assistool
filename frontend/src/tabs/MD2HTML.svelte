<script>
    import {marked} from 'marked';
    import {toastSuccess} from '../services/Toast';

    let markdownText = '';
    let htmlOutput = '';

    $: htmlOutput = marked.parse(markdownText, {async: false});

    function clearMarkdown() {
        markdownText = '';
    }

    function downloadHTML() {
        const blob = new Blob([htmlOutput], {type: 'text/html'});
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'output.html';
        a.click();
        URL.revokeObjectURL(url);
        toastSuccess('HTML downloaded!', true);
    }
</script>

<div class="row">
    <div class="col-md-6">
        <div class="card">
            <div class="card-header d-flex justify-content-between align-items-center">
                Markdown
                <button type="button" class="btn btn-outline-secondary btn-sm" on:click={clearMarkdown}>Clear</button>
            </div>
            <div class="card-body p-0">
                <textarea
                    class="form-control md-input mono"
                    bind:value={markdownText}
                    placeholder="Enter Markdown here..."
                    rows="20"
                ></textarea>
            </div>
        </div>
    </div>
    <div class="col-md-6">
        <div class="card">
            <div class="card-header d-flex justify-content-between align-items-center">
                HTML
                <button type="button" class="btn btn-outline-secondary btn-sm" on:click={downloadHTML}>
                    Download HTML
                </button>
            </div>
            <div class="card-body html-output">
                {#if htmlOutput}
                    <div>{@html htmlOutput}</div>
                {:else}
                    <p class="text-muted">HTML preview will appear here...</p>
                {/if}
            </div>
        </div>
    </div>
</div>

<style>
    .md-input {
        height: calc(100vh - 140px);
        min-height: calc(100vh - 140px);
        max-height: calc(100vh - 140px);
    }

    .html-output {
        height: calc(100vh - 140px);
        min-height: calc(100vh - 140px);
        max-height: calc(100vh - 140px);
        overflow-y: auto;
    }
</style>
