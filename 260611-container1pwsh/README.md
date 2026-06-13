Приклад контейнеру, що включає **обмеження CPU та RAM** для процесу PowerShell на Windows 11. Для цього в Go використовується **Windows Job Objects** — спеціальний механізм ОС, який дозволяє групувати процеси та накладати на них ресурсні ліміти.
---

## Ключові моменти
- **Windows Job Objects**: дозволяють задати квоти на пам’ять і процесор.
- **syscall**: у Go можна викликати Windows API напряму.
- **CreateJobObject** та **SetInformationJobObject**: основні функції для створення та налаштування обмежень.

## Що тут відбувається
- Створюється **Job Object**.
- PowerShell процес прив’язується до нього.
- Встановлюється ліміт пам’яті (100 MB).
JOBOBJECT_CPU_RATE_CONTROL_INFORMATION -- структура для керування використанням CPU.

SetInformationJobObject: використовується для застосування лімітів.

Прапори:
    JOBOBJECT_CPU_RATE_CONTROL_ENABLE (0x1) — увімкнути контроль.
    JOBOBJECT_CPU_RATE_CONTROL_HARD_CAP (0x4) — жорстке обмеження.

Значення CpuRate задається у відсотках *100 (тобто 30% → 3000).