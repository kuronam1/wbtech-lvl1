Проблемы:
1. Утечка памяти. Удержание всего среза в пати путем создания ссылки на него
2. Проблемы с возможными Unicode символами. Когда берем срез, мы можем обрезать символы, которые не помщеются в один байт (Например: символы кириллицы, которые занимают 2 байта, или символы китайского алфавита, которые могут превосходить 2 байта)