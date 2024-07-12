# Uy vazifasi: User Ranking System (Foydalanuvchilar reytingi tizimi)

## Maqsad
Foydalanuvchilar ball to'plashi mumkin bo'lgan foydalanuvchi reyting tizimini yaratish va ularning reytinglari `Redis`-da boshqariladi. Tizim quyidagilarni hal qilishi kerak:

- `Redis` set lari yordamida foydalanuvchilarni va ularning ballarini boshqarish.
- `Redis` sorted set lari yordamida saralangan peshqadamlar jadvalini saqlash.
- `Redis` tranzaktsiyalari yordamida foydalanuvchi ballari va reytinglariga atomik yangilanishlarni ta'minlash.

## Vazifalarni taqsimlash
## 1-qism: Foydalanuvchi ballarini boshqarish
1. **Add User Scores**:
- Redis sorted set ga tasodifiy ballga ega bo'lgan ko'p sonli foydalanuvchilarni qo'shing.
- Peshqadamlar jadvaliga foydalanuvchi ballarini qo'shish uchun funksiyalarni amalga oshiring.

## 2-qism: Peshqadamlar jadvalini ko'rsatish
2. **Display Leaderboard**:
- Peshqadamlar jadvalida eng yaxshi `N` foydalanuvchini oling va ko'rsating.
- Foydalanuvchilarni ballari asosida olish va ko‘rsatish funksiyalarini amalga oshiring.

## 3-qism: Tranzaksiyalar bilan ballarni yangilash
3. **Atomic Score Updates**:
- Foydalanuvchi ballarini atomik yangilash uchun `Redis` tranzaktsiyalaridan foydalaning.
- Peshqadamlar jadvali yangilangan ballarni toʻgʻri aks ettirishiga ishonch hosil qiling.
