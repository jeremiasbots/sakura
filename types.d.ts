declare module "sakura:ev3" {
    function MoveLargeMotor(port: number, speed: number, seconds: number): void;
}

interface ConsoleObject {
    log(...args: unknown[]): void;
    assert(...args: unknown[]): void;
}

declare const console: ConsoleObject;