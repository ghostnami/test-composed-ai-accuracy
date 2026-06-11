// hunk: @@ -28,7 +28,7 @@ export function initViewedCheckboxListenerFor() {
    // To prevent double addition of listeners
    form.setAttribute('data-has-viewed-checkbox-listener', String(true));

    // The checkbox consists of a div containing the real checkbox with its label,
    // hence the actual checkbox first has to be found
    const checkbox = form.querySelector<HTMLInputElement>('input[type=checkbox]')!;
    checkbox.addEventListener('input', function() {
