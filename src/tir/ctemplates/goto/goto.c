int simple_goto() {
    int x = 13;
label:
    x += 1;
    goto label;
}
