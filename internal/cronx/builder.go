package cronx

import (
	"fmt"
	"strconv"
	"strings"
)

// 星期常量定义
const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

// JobBuilder 用于构建带有执行周期的任务
type JobBuilder struct {
	job        Job
	expression string // cron 表达式
}

// 基础时间调度方法
func (jb *JobBuilder) spliceIntoPosition(position int, value string) *JobBuilder {
	segments := strings.Split(jb.expression, " ")
	if len(segments) != 5 {
		segments = []string{"*", "*", "*", "*", "*"}
	}
	segments[position-1] = value
	jb.expression = strings.Join(segments, " ")
	fmt.Println(jb.expression)
	return jb
}

// handleJob 用于执行任务
func (jb *JobBuilder) handleJob() {
	AddJob(jb.expression, jb.job)
}

// Cron 设置自定义cron表达式
func (jb *JobBuilder) Cron(expression string) {
	jb.expression = expression
	jb.handleJob()
}

// 分钟级别的调度
func (jb *JobBuilder) EveryMinute() {
	jb.spliceIntoPosition(1, "*").
		handleJob()
}

func (jb *JobBuilder) EveryTwoMinutes() {
	jb.spliceIntoPosition(1, "*/2").
		handleJob()
}

func (jb *JobBuilder) EveryThreeMinutes() {
	jb.spliceIntoPosition(1, "*/3").
		handleJob()
}

func (jb *JobBuilder) EveryFourMinutes() {
	jb.spliceIntoPosition(1, "*/4").
		handleJob()
}

func (jb *JobBuilder) EveryFiveMinutes() {
	jb.spliceIntoPosition(1, "*/5").
		handleJob()
}

func (jb *JobBuilder) EveryTenMinutes() {
	jb.spliceIntoPosition(1, "*/10").
		handleJob()
}

func (jb *JobBuilder) EveryFifteenMinutes() {
	jb.spliceIntoPosition(1, "*/15").
		handleJob()
}

func (jb *JobBuilder) EveryThirtyMinutes() {
	jb.spliceIntoPosition(1, "0,30").
		handleJob()
}

// 小时级别的调度
func (jb *JobBuilder) Hourly() {
	jb.spliceIntoPosition(1, "0").
		spliceIntoPosition(2, "*").
		handleJob()
}

func (jb *JobBuilder) HourlyAt(minute int) {
	jb.spliceIntoPosition(1, strconv.Itoa(minute)).
		handleJob()
}

func (jb *JobBuilder) EveryTwoHours() {
	jb.spliceIntoPosition(2, "*/2").
		handleJob()
}

func (jb *JobBuilder) EveryThreeHours() {
	jb.spliceIntoPosition(2, "*/3").
		handleJob()
}

func (jb *JobBuilder) EveryFourHours() {
	jb.spliceIntoPosition(2, "*/4").
		handleJob()
}

func (jb *JobBuilder) EverySixHours() {
	jb.spliceIntoPosition(2, "*/6").
		handleJob()
}

// 天级别的调度
func (jb *JobBuilder) Daily() {
	jb.spliceIntoPosition(1, "0").
		spliceIntoPosition(2, "0").
		handleJob()
}

func (jb *JobBuilder) DailyAt(time string) {
	segments := strings.Split(time, ":")
	if len(segments) != 2 {
		panic("Time should be in format HH:mm")
	}

	// 添加时间验证
	hour, err := strconv.Atoi(segments[0])
	if err != nil || hour < 0 || hour > 23 {
		panic("Hour must be between 0 and 23")
	}

	minute, err := strconv.Atoi(segments[1])
	if err != nil || minute < 0 || minute > 59 {
		panic("Minute must be between 0 and 59")
	}

	jb.spliceIntoPosition(1, segments[1]).
		spliceIntoPosition(2, segments[0]).
		handleJob()
}

// 星期几的调度
func (jb *JobBuilder) Weekdays() {
	jb.Days(fmt.Sprintf("%d-%d", Monday, Friday))
}

func (jb *JobBuilder) Weekends() {
	jb.Days(fmt.Sprintf("%d,%d"))
}

func (jb *JobBuilder) Mondays() {
	jb.Days(strconv.Itoa(Monday))
}

func (jb *JobBuilder) Tuesdays() {
	jb.Days(strconv.Itoa(Tuesday))
}

func (jb *JobBuilder) Wednesdays() {
	jb.Days(strconv.Itoa(Wednesday))
}

func (jb *JobBuilder) Thursdays() {
	jb.Days(strconv.Itoa(Thursday))
}

func (jb *JobBuilder) Fridays() {
	jb.Days(strconv.Itoa(Friday))
}

func (jb *JobBuilder) Saturdays() {
	jb.Days(strconv.Itoa(Saturday))
}

func (jb *JobBuilder) Sundays() {
	jb.Days(strconv.Itoa(Sunday))
}

func (jb *JobBuilder) Days(days string) {
	jb.spliceIntoPosition(5, days).
		handleJob()
}

// 月份相关的调度
func (jb *JobBuilder) Monthly() {
	jb.spliceIntoPosition(1, "0").
		spliceIntoPosition(2, "0").
		spliceIntoPosition(3, "1").
		handleJob()
}

func (jb *JobBuilder) MonthlyOn(dayOfMonth int, time string) {
	jb.DailyAt(time)
	jb.spliceIntoPosition(3, strconv.Itoa(dayOfMonth)).
		handleJob()
}

// 每周相关的调度
func (jb *JobBuilder) Weekly() {
	jb.spliceIntoPosition(1, "0").
		spliceIntoPosition(2, "0").
		spliceIntoPosition(5, "0").
		handleJob()
}

func (jb *JobBuilder) WeeklyOn(dayOfWeek int, time string) {
	jb.DailyAt(time)
	jb.spliceIntoPosition(5, strconv.Itoa(dayOfWeek)).
		handleJob()
}

// 每月两次
func (jb *JobBuilder) TwiceMonthly(first, second int, time string) {
	jb.DailyAt(time)
	jb.spliceIntoPosition(3, fmt.Sprintf("%d,%d", first, second)).
		handleJob()
}

// 每天两次
func (jb *JobBuilder) TwiceDaily(first, second int) {
	jb.spliceIntoPosition(2, fmt.Sprintf("%d,%d")).
		handleJob()
}

func (jb *JobBuilder) TwiceDailyAt(first, second, minute int) {
	jb.spliceIntoPosition(1, strconv.Itoa(minute)).
		spliceIntoPosition(2, fmt.Sprintf("%d,%d", first, second)).
		handleJob()
}

// 秒级调度
func (jb *JobBuilder) EverySecond() {
	jb.Cron("@every 1s")
}

func (jb *JobBuilder) EveryTwoSeconds() {
	jb.Cron("@every 2s")
}

func (jb *JobBuilder) EveryFiveSeconds() {
	jb.Cron("@every 5s")
}

func (jb *JobBuilder) EveryTenSeconds() {
	jb.Cron("@every 10s")
}

func (jb *JobBuilder) EveryFifteenSeconds() {
	jb.Cron("@every 15s")
}

func (jb *JobBuilder) EveryThirtySeconds() {
	jb.Cron("@every 30s")
}

// 季度和年度调度
func (jb *JobBuilder) Quarterly() {
	jb.spliceIntoPosition(1, "0").
		spliceIntoPosition(2, "0").
		spliceIntoPosition(3, "1").
		spliceIntoPosition(4, "1,4,7,10").
		handleJob()
}

func (jb *JobBuilder) QuarterlyOn(dayOfQuarter int, time string) {
	jb.DailyAt(time)
	jb.spliceIntoPosition(3, strconv.Itoa(dayOfQuarter)).
		spliceIntoPosition(4, "1,4,7,10").
		handleJob()
}

func (jb *JobBuilder) Yearly() {
	jb.spliceIntoPosition(1, "0").
		spliceIntoPosition(2, "0").
		spliceIntoPosition(3, "1").
		spliceIntoPosition(4, "1").
		handleJob()
}

func (jb *JobBuilder) YearlyOn(month int, dayOfMonth int, time string) {
	jb.DailyAt(time)
	jb.spliceIntoPosition(3, strconv.Itoa(dayOfMonth)).
		spliceIntoPosition(4, strconv.Itoa(month)).
		handleJob()
}

// 月末调度
func (jb *JobBuilder) LastDayOfMonth(time string) {
	segments := strings.Split(time, ":")
	if len(segments) != 2 {
		panic("Time should be in format HH:mm")
	}

	// 添加时间验证
	hour, err := strconv.Atoi(segments[0])
	if err != nil || hour < 0 || hour > 23 {
		panic("Hour must be between 0 and 23")
	}

	minute, err := strconv.Atoi(segments[1])
	if err != nil || minute < 0 || minute > 59 {
		panic("Minute must be between 0 and 59")
	}

	jb.spliceIntoPosition(1, segments[1]). // 分钟
						spliceIntoPosition(2, segments[0]). // 小时
						spliceIntoPosition(3, "L").         // 日期（月末）
						spliceIntoPosition(4, "*").         // 月份
						spliceIntoPosition(5, "*").         // 星期
						handleJob()
}

// 奇数小时调度
func (jb *JobBuilder) EveryOddHour(minute int) {
	jb.spliceIntoPosition(1, strconv.Itoa(minute)).
		spliceIntoPosition(2, "1-23/2").
		handleJob()
}

// 添加构造函数
func NewJobBuilder(job Job) *JobBuilder {
	return &JobBuilder{
		job:        job,
		expression: "* * * * *",
	}
}
